package repository

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"

	"github.com/spashev/superapp/internal/models"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repository *ProductRepository) GetAllProducts(userId, limit, offset int) (*models.ProductsPaginate, error) {
	imageBaseUrl := os.Getenv("IMAGE_BASE_URL")
	if imageBaseUrl == "" {
		return nil, fmt.Errorf("IMAGE_BASE_URL not set")
	}

	query := `
		SELECT 
			COUNT(*) OVER () AS total_count,
			p.id AS product_id,
			p.slug,
			u.id AS owner_id,
			u.email AS owner_email,
			u.first_name AS owner_first_name,
			u.last_name AS owner_last_name,
			COALESCE(u.middle_name, '') AS owner_middle_name,
			COALESCE(u.phone_number, '') AS owner_phone_number,
			COALESCE(u.avatar, '') AS owner_avatar,
			p.name_ru AS name,
			p.price_per_night,
			co.name_ru AS country,
			ci.name_ru AS city,
			p.district_ru AS district,
			p.address_ru AS address,
			CASE 
				WHEN p.created_at >= NOW() - INTERVAL '10 days' THEN true 
				ELSE false 
			END AS is_new,
			COALESCE(l.like_count, 0) AS rating,
			p.best_product,
			p.promotion,
			p.is_active,
			COALESCE(f.is_favorite, false) AS is_favorite,
			COALESCE(
				json_agg(
					DISTINCT jsonb_build_object(
						'id', pi.id,
						'thumbnail', pi.thumbnail,
						'original', pi.original,
						'mimetype', pi.mimetype,
						'is_label', pi.is_label,
						'width', pi.width,
						'height', pi.height
					)
				) FILTER (WHERE pi.id IS NOT NULL), '[]'
			) AS images
		FROM products p
		LEFT JOIN users u ON p.owner_id = u.id
		LEFT JOIN country co ON p.country_id = co.id
		LEFT JOIN city ci ON p.city_id = ci.id
		LEFT JOIN (
			SELECT product_id, COUNT(*) AS like_count
			FROM likes
			GROUP BY product_id
		) l ON p.id = l.product_id
		LEFT JOIN (
			SELECT product_id, true AS is_favorite
			FROM likes
			WHERE user_id = :userId
		) f ON p.id = f.product_id
		LEFT JOIN images pi ON p.id = pi.product_id
		GROUP BY p.id, u.id, co.id, ci.id, l.like_count, f.is_favorite
		ORDER BY p.id ASC
		LIMIT :limit OFFSET :offset;
	`

	params := map[string]interface{}{
		"userId": userId,
		"limit":  limit,
		"offset": offset,
	}

	rows, err := repository.db.NamedQuery(query, params)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var products []models.Products
	var totalCount int

	for rows.Next() {
		var temp struct {
			models.Products
			Images     string `json:"images"`
			TotalCount int    `json:"total_count"`
			IsFavorite bool   `json:"is_favorite"`
		}

		if err := rows.Scan(
			&temp.TotalCount,
			&temp.Id, &temp.Slug,
			&temp.Owner.Id, &temp.Owner.Email, &temp.Owner.First_name,
			&temp.Owner.Last_name, &temp.Owner.Middle_name, &temp.Owner.Phone_number,
			&temp.Owner.Avatar, &temp.Name, &temp.Price_per_night,
			&temp.Country, &temp.City, &temp.District,
			&temp.Address, &temp.Is_new, &temp.Rating,
			&temp.Best_product, &temp.Promotion, &temp.Is_active,
			&temp.IsFavorite, &temp.Images,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		// Устанавливаем общее количество, если это первая строка
		if totalCount == 0 {
			totalCount = temp.TotalCount
		}

		// Декодируем JSON изображений
		var images []models.ProductImages
		if err := json.Unmarshal([]byte(temp.Images), &images); err != nil {
			fmt.Println("failed to unmarshal images JSON:", err)
			continue
		}

		// Добавляем BASE_URL к изображениям
		for i := range images {
			images[i].Original = imageBaseUrl + "/" + images[i].Original
			images[i].Thumbnail = imageBaseUrl + "/" + images[i].Thumbnail
		}

		temp.Products.Images = images
		temp.Products.Is_favorite = temp.IsFavorite
		products = append(products, temp.Products)
	}

	// Рассчитываем страницы
	totalPages := (totalCount + limit - 1) / limit
	baseURL := os.Getenv("BASE_URL")

	next := ""
	if offset+limit < totalCount {
		next = fmt.Sprintf("%s/products?limit=%d&offset=%d", baseURL, limit, offset+limit)
	}

	previous := ""
	if offset > 0 {
		prevOffset := offset - limit
		if prevOffset < 0 {
			prevOffset = 0
		}
		previous = fmt.Sprintf("%s/products?limit=%d&offset=%d", baseURL, limit, prevOffset)
	}

	return &models.ProductsPaginate{
		Count:      totalCount,
		Results:    products,
		Next:       next,
		Previous:   previous,
		TotalPages: totalPages,
	}, nil
}

func (repository *ProductRepository) GetUserFavoriteProducts(userId, limit, offset int) (*models.ProductsPaginate, error) {
	imageBaseUrl := os.Getenv("IMAGE_BASE_URL")
	if imageBaseUrl == "" {
		return nil, fmt.Errorf("IMAGE_BASE_URL not set")
	}

	var totalCount int
	err := repository.db.Get(&totalCount, "SELECT COUNT(*) FROM favorites WHERE user_id = $1", userId)
	if err != nil {
		return nil, fmt.Errorf("failed to get total count: %w", err)
	}

	query := `
		SELECT 
			p.id AS product_id,
			p.slug,
			u.id AS owner_id,
			u.email AS owner_email,
			u.first_name AS owner_first_name,
			u.last_name AS owner_last_name,
			COALESCE(u.middle_name, '') AS owner_middle_name,
			COALESCE(u.phone_number, '') AS owner_phone_number,
			COALESCE(u.avatar, '') AS owner_avatar,
			p.name_ru AS name,
			p.price_per_night,
			co.name_ru AS country,
			ci.name_ru AS city,
			p.district_ru AS district,
			p.address_ru AS address,
			p.created_at >= NOW() - INTERVAL '10 days' AS is_new,
			COALESCE(l.like_count, 0) AS rating,
			p.best_product,
			p.promotion,
			p.is_active,
			EXISTS (
				SELECT 1 FROM likes WHERE product_id = p.id AND user_id = :userId
			) AS is_favorite
		FROM products p
		INNER JOIN favorites f ON f.product_id = p.id
		LEFT JOIN users u ON p.owner_id = u.id
		LEFT JOIN country co ON p.country_id = co.id
		LEFT JOIN city ci ON p.city_id = ci.id
		LEFT JOIN (
			SELECT product_id, COUNT(*) AS like_count FROM likes GROUP BY product_id
		) l ON p.id = l.product_id
		WHERE f.user_id = :userId
		ORDER BY p.id ASC
		LIMIT :limit OFFSET :offset;
	`

	params := map[string]interface{}{
		"userId": userId,
		"limit":  limit,
		"offset": offset,
	}

	rows, err := repository.db.NamedQuery(query, params)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	products := make([]models.Products, 0, limit)

	for rows.Next() {
		var product models.Products
		if err := rows.Scan(
			&product.Id, &product.Slug,
			&product.Owner.Id, &product.Owner.Email, &product.Owner.First_name,
			&product.Owner.Last_name, &product.Owner.Middle_name, &product.Owner.Phone_number,
			&product.Owner.Avatar, &product.Name, &product.Price_per_night,
			&product.Country, &product.City, &product.District,
			&product.Address, &product.Is_new, &product.Rating,
			&product.Best_product, &product.Promotion, &product.Is_active,
			&product.Is_favorite,
		); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		var images []models.ProductImages
		imageQuery := `
			SELECT id, thumbnail, original, mimetype, is_label, width, height 
			FROM images WHERE product_id = $1
		`
		err := repository.db.Select(&images, imageQuery, product.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch images: %w", err)
		}

		for i := range images {
			images[i].Original = imageBaseUrl + "/" + images[i].Original
			images[i].Thumbnail = imageBaseUrl + "/" + images[i].Thumbnail
		}

		product.Images = images
		products = append(products, product)
	}

	totalPages := (totalCount + limit - 1) / limit
	baseURL := os.Getenv("BASE_URL")

	var next, previous string
	if offset+limit < totalCount {
		next = fmt.Sprintf("%s/products?limit=%d&offset=%d", baseURL, limit, offset+limit)
	}
	if offset > 0 {
		prevOffset := max(offset-limit, 0)
		previous = fmt.Sprintf("%s/products?limit=%d&offset=%d", baseURL, limit, prevOffset)
	}

	return &models.ProductsPaginate{
		Count:      totalCount,
		Results:    products,
		Next:       next,
		Previous:   previous,
		TotalPages: totalPages,
	}, nil
}

func (repository *ProductRepository) GetProductBySlug(slug string) (*models.Product, error) {
	query := `
		SELECT 
			p.id,
			p.slug,
			p.name_ru AS name,
			p.price_per_night,
			p.price_per_week,
			p.price_per_month,
			u.id AS "owner.id",
			u.email AS "owner.email",
			u.first_name AS "owner.first_name",
			u.last_name AS "owner.last_name",
			u.middle_name AS "owner.middle_name",
			u.phone_number AS "owner.phone_number",
			u.avatar AS "owner.avatar",
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
			p.lng,
			p.lat,
			COALESCE(AVG(l.like_count), 0) AS average_likes_rating,
			p.phone_number,
			CASE
				WHEN p.created_at >= NOW() - INTERVAL '10 days' THEN true
				ELSE false
			END AS is_new,
			p.best_product,
			p.promotion,
			p.type_id
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
	`

	var product models.Product
	err := repository.db.QueryRowx(query, slug).StructScan(&product)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch product by slug: %w", err)
	}

	product.Images, _ = repository.getImagesByProductID(product.Id)
	product.Comments, _ = repository.getCommentsByProductId(product.Id)
	productType, _ := repository.getProductTypeById(product.Type_id)
	if productType != nil {
		product.Type = *productType
	}
	product.Conveniences, _ = repository.getConveniencesByProductId(product.Id)

	return &product, nil
}

func (repository *ProductRepository) LikeProductById(userID, productID int) error {
	tx, err := repository.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var likeID int
	checkLikeQuery := `
		SELECT id FROM likes WHERE product_id = $1 AND user_id = $2
	`
	err = tx.Get(&likeID, checkLikeQuery, productID, userID)

	if err == nil {
		return nil
	} else if errors.Is(err, sql.ErrNoRows) {
		insertLikeQuery := `
			INSERT INTO likes (product_id, user_id, count, created_at, updated_at) 
			VALUES ($1, $2, 1, $3, $4) RETURNING id
		`
		err = tx.QueryRow(insertLikeQuery, productID, userID, time.Now(), time.Now()).Scan(&likeID)
		if err != nil {
			return err
		}

		insertFavoriteQuery := `
			INSERT INTO favorites (like_id, product_id, user_id, created_at, updated_at) 
			VALUES ($1, $2, $3, $4, $5)
		`
		_, err = tx.Exec(insertFavoriteQuery, likeID, productID, userID, time.Now(), time.Now())
		if err != nil {
			return err
		}

		updateQuery := `UPDATE products SET like_count = like_count + 1 WHERE id = $1`
		_, err = tx.Exec(updateQuery, productID)
		if err != nil {
			return err
		}
	} else {
		return err
	}

	return nil
}

func (repository *ProductRepository) DislikeProductById(userID, productID int) error {
	tx, err := repository.db.Beginx()
	if err != nil {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	var likeID int
	checkLikeQuery := `
		SELECT id FROM likes WHERE product_id = $1 AND user_id = $2
	`
	err = tx.Get(&likeID, checkLikeQuery, productID, userID)

	if err == nil {
		deleteLikeQuery := `DELETE FROM likes WHERE id = $1`
		_, err = tx.Exec(deleteLikeQuery, likeID)
		if err != nil {
			return err
		}

		deleteFavoriteQuery := `DELETE FROM favorites WHERE like_id = $1`
		_, err = tx.Exec(deleteFavoriteQuery, likeID)
		if err != nil {
			return err
		}

		updateQuery := `UPDATE products SET like_count = like_count - 1 WHERE id = $1`
		_, err = tx.Exec(updateQuery, productID)
		if err != nil {
			return err
		}
	} else if errors.Is(err, sql.ErrNoRows) {
		return nil
	} else {
		return err
	}

	return nil
}

func (repository *ProductRepository) getImagesByProductID(productID int) ([]models.ProductImages, error) {
	imageBaseUrl := os.Getenv("IMAGE_BASE_URL")
	if imageBaseUrl == "" {
		return nil, fmt.Errorf("IMAGE_BASE_URL not set")
	}

	query := `
		SELECT id, original, thumbnail, mimetype, is_label, width, height
		FROM images
		WHERE product_id = $1;
	`

	var images []models.ProductImages
	err := repository.db.Select(&images, query, productID)
	if err != nil {
		return nil, err
	}
	for i := range images {
		images[i].Original = imageBaseUrl + "/" + images[i].Original
		images[i].Thumbnail = imageBaseUrl + "/" + images[i].Thumbnail
	}

	return images, nil
}

func (repository *ProductRepository) getCommentsByProductId(product_id int) ([]models.ProductComment, error) {
	query := `
		SELECT 
			c.id, 
			c.content,
			c.rating, 
			c.created_at,
			u.id AS "user.id",
			u.email AS "user.email",
			u.first_name AS "user.first_name",
			u.last_name AS "user.last_name",
			u.avatar AS "user.avatar"
		FROM 
			comments AS c
		LEFT JOIN users AS u ON u.id = c.user_id
		WHERE 
			c.product_id = $1;
	`

	var comments []models.ProductComment
	err := repository.db.Select(&comments, query, product_id)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (repository *ProductRepository) getProductTypeById(product_id int) (*models.ProductType, error) {
	query := `
		SELECT 
			id,
			icon,
			name_ru AS name
		FROM types
		WHERE id = $1;
	`
	var productType models.ProductType
	err := repository.db.Get(&productType, query, product_id)
	if err != nil {
		return nil, err
	}

	return &productType, nil
}

func (repository *ProductRepository) getConveniencesByProductId(product_id int) ([]models.Convenience, error) {
	imageBaseUrl := os.Getenv("IMAGE_BASE_URL")
	if imageBaseUrl == "" {
		return nil, fmt.Errorf("IMAGE_BASE_URL not set")
	}

	query := `
		SELECT 
			c.id,
			c.icon,
			c.slug,
			c.name_ru AS name
		FROM products_convenience pc
		INNER JOIN conveniences c ON c.id = pc.convenience_id
		WHERE pc.product_id = $1;
	`

	var conveniences []models.Convenience
	err := repository.db.Select(&conveniences, query, product_id)
	if err != nil {
		return nil, err
	}

	for i := range conveniences {
		conveniences[i].Icon = imageBaseUrl + "/" + conveniences[i].Icon
	}

	return conveniences, nil
}
