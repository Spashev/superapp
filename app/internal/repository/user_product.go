package repository

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"

	"github.com/spashev/superapp/internal/models"
)

type UserProductRepository struct {
	db *sqlx.DB
}

func NewUserProductRepository(db *sqlx.DB) *UserProductRepository {
	return &UserProductRepository{db: db}
}

func (repository *UserProductRepository) GetUserFavoriteProducts(userId, limit, offset int) (*models.ProductsPaginate, error) {
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
