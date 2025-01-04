package repository

import (
	"database/sql"
	"sync"

	"superapp/models"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAllProducts(page, limit int) ([]models.ProductPaginate, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Query(`
		SELECT 
			p.id AS product_id,
			p.slug,
			p.name_ru AS product_name,
			p.price_per_night,
			u.email AS email,
			u.first_name AS first_name,
			u.last_name AS last_name,
			co.name_ru AS country_name,
			ci.name_ru AS city_name,
			p.district_ru,
			p.address_ru,
			CASE
				WHEN p.created_at >= NOW() - INTERVAL '10 days' THEN true
				ELSE false
			END AS is_new,
			COALESCE(AVG(l.like_count), 0) AS rating,
			p.best_product,
			p.promotion,
			p.is_active,
			img.id AS image_id,
			img.thumbnail,
			img.mimetype,
			img.is_label,
			img.width,
			img.height,
			p.created_at,
			p.updated_at
		FROM 
			products p
		LEFT JOIN users u ON p.owner_id = u.id
		LEFT JOIN country co ON p.country_id = co.id
		LEFT JOIN city ci ON p.city_id = ci.id
		LEFT JOIN (
			SELECT product_id, COUNT(*) AS like_count
			FROM likes
			GROUP BY product_id
		) l ON p.id = l.product_id
		LEFT JOIN images img ON p.id = img.product_id
		GROUP BY 
			p.id, 
			u.id, 
			co.id,
			ci.id,
			img.id
		ORDER BY p.created_at
		LIMIT $1 OFFSET $2;
	`, limit, offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	productsChan := make(chan models.ProductPaginate, limit)
	baseURL := "http://localhost:9001/"

	var wg sync.WaitGroup
	for rows.Next() {
		var p models.ProductPaginate
		var image models.ProductImagePaginate
		var owner models.Owner

		if err := rows.Scan(
			&p.ID, &p.Slug, &p.NameRU, &p.PricePerNight, &owner.Email, &owner.FirstName, &owner.LastName,
			&p.CountryName, &p.CityName, &p.DistrictRU, &p.AddressRU, &p.IsNew, &p.Rating, &p.BestProduct,
			&p.Promotion, &p.IsActive, &image.ID, &image.Thumbnail, &image.MimeType, &image.IsLabel, &image.Width,
			&image.Height, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}

		wg.Add(1)

		go func(productID int64, image models.ProductImagePaginate, owner models.Owner) {
			defer wg.Done()

			if image.Thumbnail != "" {
				image.Thumbnail = baseURL + image.Thumbnail
			}

			product := models.ProductPaginate{
				ID:            p.ID,
				Slug:          p.Slug,
				NameRU:        p.NameRU,
				PricePerNight: p.PricePerNight,
				CountryName:   p.CountryName,
				CityName:      p.CityName,
				Owner:         owner,
				IsNew:         p.IsNew,
				Rating:        p.Rating,
				BestProduct:   p.BestProduct,
				Promotion:     p.Promotion,
				IsActive:      p.IsActive,
				Images:        []models.ProductImagePaginate{image},
			}

			productsChan <- product
		}(p.ID, image, owner)
	}

	go func() {
		wg.Wait()
		close(productsChan)
	}()

	var products []models.ProductPaginate
	for product := range productsChan {
		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) GetProductBySlug(slug string) (*models.Product, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	var p models.Product
	var owner models.OwnerProduct
	var productType models.ProductType
	var comments []models.ProductComment
	var conveniences []models.Convenience
	var images []models.ProductImagePaginate
	var err error

	wg.Add(1)
	go func() {
		defer wg.Done()
		pRows, queryErr := r.db.Query(`
			SELECT 
				p.id,
				p.slug,
				p.name_ru AS name,
				p.price_per_night,
				p.price_per_week,
				p.price_per_month,
				u.id AS user_id,
				u.email AS email,
				u.first_name,
				u.last_name,
				u.middle_name,
				u.phone_number AS phone_number,
				u.avatar,
				p.rooms_qty,
				p.guest_qty,
				p.bed_qty,
				p.bedroom_qty,
				p.toilet_qty,
				p.bath_qty,
				p.description_ru AS description,
				co.name_ru AS country,
				ci.name_ru AS city,
				p.district_ru AS district,
				p.address_ru AS address,
				p.like_count,
				COALESCE(AVG(l.like_count), 0) AS average_likes_rating,
				p.phone_number,
				CASE
					WHEN p.created_at >= NOW() - INTERVAL '10 days' THEN true
					ELSE false
				END AS is_new,
				p.best_product,
				p.promotion,
				img.id AS image_id,
				img.thumbnail,
				img.mimetype,
				img.is_label,
				img.width,
				img.height
			FROM 
				products p
			LEFT JOIN users u ON p.owner_id = u.id
			LEFT JOIN country co ON p.country_id = co.id
			LEFT JOIN city ci ON p.city_id = ci.id
			LEFT JOIN (
				SELECT product_id, COUNT(*) AS like_count
				FROM likes
				GROUP BY product_id
			) l ON p.id = l.product_id
			LEFT JOIN images img ON p.id = img.product_id
			WHERE p.slug = $1
			GROUP BY 
				p.id, 
				u.id, 
				co.id,
				ci.id,
				img.id
			ORDER BY p.created_at;
		`, slug)

		if queryErr != nil {
			err = queryErr
			return
		}
		defer pRows.Close()

		for pRows.Next() {
			var image models.ProductImagePaginate
			if scanErr := pRows.Scan(
				&p.ID, &p.Slug, &p.Name, &p.PricePerNight, &p.PricePerWeek, &p.PricePerMonth,
				&owner.ID, &owner.Email, &owner.FirstName, &owner.LastName, &owner.MiddleName, &owner.PhoneNumber, &owner.Avatar,
				&p.RoomsQty, &p.GuestQty, &p.BedQty, &p.BedroomQty, &p.ToiletQty, &p.BathQty, &p.Description,
				&p.Country, &p.City, &p.District, &p.Address, &p.LikeCount, &p.AverageLikesRating,
				&p.PhoneNumber, &p.IsNew, &p.BestProduct, &p.Promotion,
				&image.ID, &image.Thumbnail, &image.MimeType, &image.IsLabel, &image.Width, &image.Height,
			); scanErr != nil {
				err = scanErr
				return
			}
			mu.Lock()
			images = append(images, image)
			mu.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		cRows, queryErr := r.db.Query(`
			SELECT 
				c.id AS comment_id,
				c.content AS comment_content,
				c.rating AS comment_rating,
				cu.email AS comment_user_email,
				cu.first_name AS comment_first_name,
				cu.last_name AS comment_last_name,
				cu.avatar AS comment_avatar
			FROM 
				comments c
			LEFT JOIN users cu ON cu.id = c.user_id
			WHERE c.product_id = (SELECT id FROM products WHERE slug = $1);
		`, slug)

		if queryErr != nil {
			err = queryErr
			return
		}
		defer cRows.Close()

		for cRows.Next() {
			var comment models.ProductComment
			var comment_user models.ProductCommentUser
			if scanErr := cRows.Scan(
				&comment.ID, &comment.Content, &comment.Rating,
				&comment_user.Email, &comment_user.FirstName, &comment_user.LastName, &comment_user.Avatar,
			); scanErr != nil {
				err = scanErr
				return
			}
			mu.Lock()
			comments = append(comments, comment)
			mu.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		tRows, queryErr := r.db.Query(`
			SELECT 
				t.name_ru AS name,
				t.icon AS icon
			FROM 
				types t
			LEFT JOIN products p ON t.id = p.type_id
			WHERE p.slug = $1;
		`, slug)

		if queryErr != nil {
			err = queryErr
			return
		}
		defer tRows.Close()

		for tRows.Next() {
			if scanErr := tRows.Scan(
				&productType.Name, &productType.Icon,
			); scanErr != nil {
				err = scanErr
				return
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		pcRows, queryErr := r.db.Query(`
			SELECT 
				conv.id AS id,
				conv.name_ru AS name,
				conv.icon AS icon
			FROM 
				products_convenience pc
			LEFT JOIN conveniences conv ON pc.convenience_id = conv.id
			LEFT JOIN products p ON p.id = pc.product_id
			WHERE p.slug = $1;
		`, slug)

		if queryErr != nil {
			err = queryErr
			return
		}
		defer pcRows.Close()

		for pcRows.Next() {
			var convenience models.Convenience
			if scanErr := pcRows.Scan(
				&convenience.ID, &convenience.Name, &convenience.Icon,
			); scanErr != nil {
				err = scanErr
				return
			}
			mu.Lock()
			conveniences = append(conveniences, convenience)
			mu.Unlock()
		}
	}()

	wg.Wait()

	if err != nil {
		return nil, err
	}

	p.Images = images
	p.Type = productType
	p.Comments = comments
	p.Conveniences = conveniences
	p.Owner = owner

	return &p, nil
}
