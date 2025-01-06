package handler

type ProductRes struct {
	ID               int64   `json:"id"`
	PricePerNight    int     `json:"price_per_night"`
	PricePerWeek     *int    `json:"price_per_week"`
	PricePerMonth    *int    `json:"price_per_month"`
	RoomsQty         int     `json:"rooms_qty"`
	GuestQty         int     `json:"guest_qty"`
	BedQty           int     `json:"bed_qty"`
	BedroomQty       int     `json:"bedroom_qty"`
	ToiletQty        *int    `json:"toilet_qty"`
	BathQty          *int    `json:"bath_qty"`
	CityID           *int64  `json:"city_id"`
	Lng              *string `json:"lng"`
	Lat              *string `json:"lat"`
	IsActive         bool    `json:"is_active"`
	Priority         string  `json:"priority"`
	LikeCount        int     `json:"like_count"`
	CommentsRu       *string `json:"comments_ru"`
	OwnerID          int64   `json:"owner_id"`
	TypeID           int64   `json:"type_id"`
	GuestsWithBabies bool    `json:"guests_with_babies"`
	GuestsWithPets   bool    `json:"guests_with_pets"`
	BestProduct      bool    `json:"best_product"`
	Promotion        bool    `json:"promotion"`
	CountryID        *int64  `json:"country_id"`
	PhoneNumber      *string `json:"phone_number"`
	AddressEn        *string `json:"address_en"`
	AddressKz        *string `json:"address_kz"`
	AddressRu        *string `json:"address_ru"`
	CommentsEn       *string `json:"comments_en"`
	CommentsKz       *string `json:"comments_kz"`
	DescriptionEn    *string `json:"description_en"`
	DescriptionKz    *string `json:"description_kz"`
	DescriptionRu    *string `json:"description_ru"`
	DistrictEn       *string `json:"district_en"`
	DistrictKz       *string `json:"district_kz"`
	DistrictRu       *string `json:"district_ru"`
	NameEn           *string `json:"name_en"`
	NameKz           *string `json:"name_kz"`
	NameRu           *string `json:"name_ru"`
	Slug             *string `json:"slug"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
}
