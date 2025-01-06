package models

import "time"

type Product struct {
	ID                 int                    `json:"id"`
	Slug               string                 `json:"slug"`
	Name               string                 `json:"name"`
	PricePerNight      int                    `json:"price_per_night"`
	PricePerWeek       int                    `json:"price_per_week"`
	PricePerMonth      int                    `json:"price_per_month"`
	Owner              OwnerProduct           `json:"owner"`
	RoomsQty           int                    `json:"rooms_qty"`
	GuestQty           int                    `json:"guest_qty"`
	BedQty             int                    `json:"bed_qty"`
	BedroomQty         int                    `json:"bedroom_qty"`
	ToiletQty          int                    `json:"toilet_qty"`
	BathQty            int                    `json:"bath_qty"`
	Description        string                 `json:"description"`
	Country            string                 `json:"country"`
	City               string                 `json:"city"`
	District           string                 `json:"district"`
	Address            string                 `json:"address"`
	LikeCount          int                    `json:"like_count"`
	AverageLikesRating float64                `json:"average_likes_rating"`
	PhoneNumber        string                 `json:"phone_number"`
	IsNew              bool                   `json:"is_new"`
	BestProduct        bool                   `json:"best_product"`
	Promotion          bool                   `json:"promotion"`
	Images             []ProductImagePaginate `json:"images"`
	Type               ProductType            `json:"type"`
	Conveniences       []Convenience          `json:"conveniences"`
	Comments           []ProductComment       `json:"comments"`
}

type ProductPaginate struct {
	ID             int64                  `json:"product_id"`
	Slug           string                 `json:"slug"`
	NameRU         string                 `json:"product_name"`
	PricePerNight  float64                `json:"price_per_night"`
	Owner          Owner                  `json:"owner"`
	OwnerFirstName string                 `json:"owner_first_name"`
	OwnerLastName  string                 `json:"owner_last_name"`
	CountryName    string                 `json:"country_name"`
	CityName       string                 `json:"city_name"`
	DistrictRU     string                 `json:"district_ru"`
	AddressRU      string                 `json:"address_ru"`
	IsNew          bool                   `json:"is_new"`
	Rating         float64                `json:"rating"`
	BestProduct    bool                   `json:"best_product"`
	Promotion      bool                   `json:"promotion"`
	IsActive       bool                   `json:"is_active"`
	Images         []ProductImagePaginate `json:"images"`
	CreatedAt      time.Time              `json:"created_at"`
	UpdatedAt      time.Time              `json:"updated_at"`
}

type ProductType struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}
