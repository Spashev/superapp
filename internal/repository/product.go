package repository

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	"superapp/internal/models"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (repository *ProductRepository) GetAllProducts(page, limit int) (*models.ProductsPaginate, error) {
	offset := (page - 1) * limit

	query := `
		SELECT 
			p.id,
			p.slug,
			u.id AS "owner.id",
			u.email AS "owner.email",
			u.first_name AS "owner.first_name",
			u.last_name AS "owner.last_name",
			u.middle_name AS "owner.middle_name",
			u.phone_number AS "owner.phone_number",
			u.avatar AS "owner.avatar",
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
			COALESCE(AVG(l.like_count), 0) AS rating,
			p.best_product,
			p.promotion,
			p.is_active,
			COUNT(*) OVER() AS total_count
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
		ORDER BY p.id ASC
		LIMIT :limit OFFSET :offset;
	`

	params := map[string]interface{}{
		"limit":  limit,
		"offset": offset,
	}

	rows, err := repository.db.NamedQuery(query, params)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	var products []models.Products
	var totalCount int64
	for rows.Next() {
		var product models.Products

		err = rows.StructScan(&product)
		if err != nil {
			return nil, fmt.Errorf("failed to scan product and owner: %w", err)
		}

		images, err := repository.getImagesByProductID(product.Id)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch images for product %d: %w", product.Id, err)
		}

		product.Images = images

		products = append(products, product)

		totalCount = product.Total_count
	}

	baseURL := os.Getenv("BASE_URL")
	next := ""
	if int64(offset+limit) < totalCount {
		next = fmt.Sprintf("%s/products/get?limit=%d&offset=%d", baseURL, limit, offset+limit)
	}
	previous := ""
	if offset > 0 {
		prevOffset := offset - limit
		if prevOffset < 0 {
			prevOffset = 0
		}
		previous = fmt.Sprintf("%s/products/get?limit=%d&offset=%d", baseURL, limit, prevOffset)
	}

	return &models.ProductsPaginate{
		Count:    totalCount,
		Results:  products,
		Next:     next,
		Previous: previous,
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

func (repository *ProductRepository) getImagesByProductID(productID int64) ([]models.ProductImages, error) {
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

func (repository *ProductRepository) getCommentsByProductId(product_id int64) ([]models.ProductComment, error) {
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

func (repository *ProductRepository) getProductTypeById(product_id int64) (*models.ProductType, error) {
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

func (repository *ProductRepository) getConveniencesByProductId(product_id int64) ([]models.Convenience, error) {
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
