package repository

import (
	"database/sql"
	"sync"

	"superapp/models"
)

type ProductRepository struct {
	db *sql.DB
	mu sync.Mutex
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAllProducts() ([]models.ProductPaginate, error) {
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
			img.id;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	productsMap := make(map[int64]*models.ProductPaginate)
	var wg sync.WaitGroup

	productsChan := make(chan *models.ProductPaginate)

	for rows.Next() {
		var p models.ProductPaginate
		var image models.ProductImagePaginate
		var owner models.Owner

		if err := rows.Scan(
			&p.ID, &p.Slug, &p.NameRU, &p.PricePerNight, &owner.Email, &owner.FirstName, &owner.LastName,
			&p.CountryName, &p.CityName, &p.DistrictRU, &p.AddressRU, &p.IsNew, &p.Rating, &p.BestProduct,
			&p.Promotion, &p.IsActive, &image.ID, &image.Thumbnail, &image.MimeType, &image.IsLabel, &image.ImageWidth,
			&image.ImageHeight, &p.CreatedAt, &p.UpdatedAt,
		); err != nil {
			return nil, err
		}

		wg.Add(1)

		go func(productID int64, image models.ProductImagePaginate, owner models.Owner) {
			defer wg.Done()

			r.mu.Lock()
			defer r.mu.Unlock()

			if _, exists := productsMap[productID]; !exists {
				productsMap[productID] = &models.ProductPaginate{
					ID:            productID,
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
				}
			}

			productsMap[productID].Images = append(productsMap[productID].Images, image)
			productsMap[productID].Owner = owner

		}(p.ID, image, owner)
	}

	go func() {
		wg.Wait()
		for _, product := range productsMap {
			productsChan <- product
		}
		close(productsChan)
	}()

	var products []models.ProductPaginate
	for product := range productsChan {
		products = append(products, *product)
	}

	return products, nil
}
