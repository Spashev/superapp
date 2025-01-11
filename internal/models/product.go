package models

import "time"

type Product struct {
	ID                 int              `json:"id"`
	Slug               string           `json:"slug"`
	Name               string           `json:"name"`
	PricePerNight      int              `json:"price_per_night"`
	PricePerWeek       int              `json:"price_per_week"`
	PricePerMonth      int              `json:"price_per_month"`
	Owner              OwnerProduct     `json:"owner"`
	RoomsQty           int              `json:"rooms_qty"`
	GuestQty           int              `json:"guest_qty"`
	BedQty             int              `json:"bed_qty"`
	BedroomQty         int              `json:"bedroom_qty"`
	ToiletQty          int              `json:"toilet_qty"`
	BathQty            int              `json:"bath_qty"`
	Description        string           `json:"description"`
	Country            string           `json:"country"`
	City               string           `json:"city"`
	District           string           `json:"district"`
	Address            string           `json:"address"`
	Conveniences       []Convenience    `json:"convenience"`
	Type               ProductType      `json:"type"`
	Lat                string           `json:"lat"`
	Lng                string           `json:"lng"`
	LikeCount          int              `json:"like_count"`
	AverageLikesRating float64          `json:"rating"`
	PhoneNumber        string           `json:"phone_number"`
	IsNew              bool             `json:"is_new"`
	BestProduct        bool             `json:"best_product"`
	Promotion          bool             `json:"promotion"`
	ProductComments    []ProductComment `json:"product_comments"`
	Comments           []string         `json:"comments"`
	Bookings           []string         `json:"bookings"`
	Images             []ProductImages  `json:"images"`
}

type Products struct {
	ID            int64           `json:"product_id"`
	Slug          string          `json:"slug"`
	NameRU        string          `json:"name"`
	PricePerNight float64         `json:"price_per_night"`
	Owner         OwnerProduct    `json:"owner"`
	CountryName   string          `json:"country"`
	CityName      string          `json:"city"`
	DistrictRU    string          `json:"district"`
	AddressRU     string          `json:"address"`
	IsNew         bool            `json:"is_new"`
	Rating        float64         `json:"rating"`
	BestProduct   bool            `json:"best_product"`
	Promotion     bool            `json:"promotion"`
	IsActive      bool            `json:"is_active"`
	Images        []ProductImages `json:"images"`
	CreatedAt     time.Time       `json:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"`
}

type ProductsPaginate struct {
	Count    int64      `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Products `json:"results"`
}

type ProductType struct {
	Name string `json:"name"`
	Icon string `json:"icon"`
}
