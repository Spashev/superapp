package handler

import (
	"encoding/json"
	"log"
	"net/http"

	store "superapp/db"
	"superapp/models"
)

func GetProductList(w http.ResponseWriter, r *http.Request) {
	dsn := "postgres://bookituser:Lenger1992@localhost:6432/bookitdb?sslmode=disable"

	database, err := store.NewDatabase(dsn)
	if err != nil {
		http.Error(w, "Error connecting to the database", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer database.Close()

	rows, err := database.Conn.Query(`
		SELECT id, price_per_night, price_per_week, price_per_month, rooms_qty, guest_qty, bed_qty, bedroom_qty, toilet_qty,
			bath_qty, city_id, lng, lat, is_active, priority, like_count, comments_ru, owner_id, type_id, guests_with_babies,
			guests_with_pets, best_product, promotion, country_id, phone_number, address_en, address_kz, address_ru,
			comments_en, comments_kz, description_en, description_kz, description_ru, district_en, district_kz, district_ru,
			name_en, name_kz, name_ru, slug, created_at, updated_at
		FROM products`)
	if err != nil {
		http.Error(w, "Error fetching products", http.StatusInternalServerError)
		log.Println(err)
		return
	}
	defer rows.Close()

	var products []models.Product
	for rows.Next() {
		var p models.Product
		if err := rows.Scan(
			&p.ID, &p.PricePerNight, &p.PricePerWeek, &p.PricePerMonth, &p.RoomsQty, &p.GuestQty, &p.BedQty, &p.BedroomQty,
			&p.ToiletQty, &p.BathQty, &p.CityID, &p.Lng, &p.Lat, &p.IsActive, &p.Priority, &p.LikeCount, &p.CommentsRu,
			&p.OwnerID, &p.TypeID, &p.GuestsWithBabies, &p.GuestsWithPets, &p.BestProduct, &p.Promotion, &p.CountryID,
			&p.PhoneNumber, &p.AddressEn, &p.AddressKz, &p.AddressRu, &p.CommentsEn, &p.CommentsKz, &p.DescriptionEn,
			&p.DescriptionKz, &p.DescriptionRu, &p.DistrictEn, &p.DistrictKz, &p.DistrictRu, &p.NameEn, &p.NameKz,
			&p.NameRu, &p.Slug, &p.CreatedAt, &p.UpdatedAt); err != nil {
			http.Error(w, "Error scanning product", http.StatusInternalServerError)
			log.Println(err)
			return
		}
		products = append(products, p)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(products); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		log.Println(err)
	}
}
