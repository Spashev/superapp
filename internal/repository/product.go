package repository

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"superapp/internal/models"
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
		GROUP BY 
			p.id, 
			u.id, 
			co.id,
			ci.id
		LIMIT $1 OFFSET $2;
	`, limit, offset)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.ProductPaginate
	var wg sync.WaitGroup
	productCh := make(chan models.ProductPaginate, limit)

	for rows.Next() {
		var product models.ProductPaginate
		var owner models.Owner

		if err := rows.Scan(
			&product.ID, &product.Slug, &product.NameRU, &product.PricePerNight, &owner.Email, &owner.FirstName,
			&owner.LastName, &product.CountryName, &product.CityName, &product.DistrictRU, &product.AddressRU,
			&product.IsNew, &product.Rating, &product.BestProduct, &product.Promotion, &product.IsActive,
			&product.CreatedAt, &product.UpdatedAt,
		); err != nil {
			return nil, err
		}

		product.Owner = owner

		wg.Add(1)
		go func(product models.ProductPaginate) {
			defer wg.Done()
			images, err := r.getImagesByProductID(product.ID)
			if err != nil {
				return
			}
			product.Images = images
			productCh <- product
		}(product)
	}

	go func() {
		wg.Wait()
		close(productCh)
	}()

	for product := range productCh {
		products = append(products, product)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductBySlug(slug string) (*models.Product, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	var p models.Product
	var owner models.OwnerProduct
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
				p.promotion
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
			WHERE p.slug = $1
			GROUP BY 
				p.id, 
				u.id, 
				co.id,
				ci.id
			ORDER BY p.created_at;
		`, slug)

		if queryErr != nil {
			err = queryErr
			return
		}
		defer pRows.Close()

		for pRows.Next() {
			if scanErr := pRows.Scan(
				&p.ID, &p.Slug, &p.Name, &p.PricePerNight, &p.PricePerWeek, &p.PricePerMonth,
				&owner.ID, &owner.Email, &owner.FirstName, &owner.LastName, &owner.MiddleName, &owner.PhoneNumber, &owner.Avatar,
				&p.RoomsQty, &p.GuestQty, &p.BedQty, &p.BedroomQty, &p.ToiletQty, &p.BathQty, &p.Description,
				&p.Country, &p.City, &p.District, &p.Address, &p.LikeCount, &p.AverageLikesRating,
				&p.PhoneNumber, &p.IsNew, &p.BestProduct, &p.Promotion,
			); scanErr != nil {
				err = scanErr
				return
			}
			mu.Lock()
			p.Images, err = r.getImagesByProductID(int64(p.ID))
			if err != nil {
				mu.Unlock()
				return
			}
			mu.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		p.Type, err = r.GetProductTypeBySlug(slug)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		p.Comments, err = r.GetCommentsByProductSlug(slug)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		p.Conveniences, err = r.GetConveniencesByProductSlug(slug)
	}()

	wg.Wait()

	if err != nil {
		return nil, err
	}

	p.Owner = owner
	return &p, nil
}

func (r *ProductRepository) getImagesByProductID(productID int64) ([]models.ProductImagePaginate, error) {
	rows, err := r.db.Query(`
		SELECT 
			id,
			thumbnail,
			mimetype,
			is_label,
			width,
			height
		FROM 
			images
		WHERE 
			product_id = $1;
	`, productID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	baseURL := os.Getenv("BASE_URL")
	if baseURL == "" {
		return nil, fmt.Errorf("BASE_URL is not set in environment")
	}

	var images []models.ProductImagePaginate

	for rows.Next() {
		var image models.ProductImagePaginate

		if err := rows.Scan(&image.ID, &image.Thumbnail, &image.MimeType,
			&image.IsLabel, &image.Width, &image.Height); err != nil {
			return nil, err
		}

		if image.Thumbnail != "" {
			image.Thumbnail = baseURL + image.Thumbnail
		}

		images = append(images, image)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return images, nil
}

func (r *ProductRepository) GetCommentsByProductSlug(slug string) ([]models.ProductComment, error) {
	var comments []models.ProductComment

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
		return nil, queryErr
	}
	defer cRows.Close()

	for cRows.Next() {
		var comment models.ProductComment
		var commentUser models.ProductCommentUser
		if scanErr := cRows.Scan(
			&comment.ID, &comment.Content, &comment.Rating,
			&commentUser.Email, &commentUser.FirstName, &commentUser.LastName, &commentUser.Avatar,
		); scanErr != nil {
			return nil, scanErr
		}
		comment.User = commentUser
		comments = append(comments, comment)
	}

	return comments, nil
}

func (r *ProductRepository) GetProductTypeBySlug(slug string) (models.ProductType, error) {
	var productType models.ProductType

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
		return productType, queryErr
	}
	defer tRows.Close()

	if tRows.Next() {
		if scanErr := tRows.Scan(&productType.Name, &productType.Icon); scanErr != nil {
			return productType, scanErr
		}
	}

	return productType, nil
}

func (r *ProductRepository) GetConveniencesByProductSlug(slug string) ([]models.Convenience, error) {
	var conveniences []models.Convenience

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
		return nil, queryErr
	}
	defer pcRows.Close()

	for pcRows.Next() {
		var convenience models.Convenience
		if scanErr := pcRows.Scan(&convenience.ID, &convenience.Name, &convenience.Icon); scanErr != nil {
			return nil, scanErr
		}
		conveniences = append(conveniences, convenience)
	}

	return conveniences, nil
}
