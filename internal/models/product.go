package models

type Product struct {
	Id                 int64            `json:"id"`
	Slug               string           `json:"slug"`
	Name               string           `json:"name"`
	PricePerNight      int64            `json:"price_per_night"`
	PricePerWeek       int64            `json:"price_per_week"`
	PricePerMonth      int64            `json:"price_per_month"`
	Owner              OwnerProduct     `json:"owner"`
	RoomsQty           int64            `json:"rooms_qty"`
	GuestQty           int64            `json:"guest_qty"`
	BedQty             int64            `json:"bed_qty"`
	BedroomQty         int64            `json:"bedroom_qty"`
	ToiletQty          int64            `json:"toilet_qty"`
	BathQty            int64            `json:"bath_qty"`
	Description        string           `json:"description"`
	Country            string           `json:"country"`
	City               string           `json:"city"`
	District           string           `json:"district"`
	Address            string           `json:"address"`
	Conveniences       []Convenience    `json:"convenience"`
	Type               ProductType      `json:"type"`
	Lat                string           `json:"lat"`
	Lng                string           `json:"lng"`
	LikeCount          int64            `json:"like_count"`
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
	Id              int64           `json:"id"`
	Slug            string          `json:"slug"`
	Name            string          `json:"name"`
	Price_per_night int64           `json:"price_per_night"`
	Owner           OwnerProduct    `json:"owner"`
	Country         string          `json:"country"`
	City            string          `json:"city"`
	District        string          `json:"district"`
	Address         string          `json:"address"`
	Is_new          bool            `json:"is_new"`
	Rating          float64         `json:"rating"`
	Best_product    bool            `json:"best_product"`
	Promotion       bool            `json:"promotion"`
	Is_active       bool            `json:"is_active"`
	Images          []ProductImages `json:"images"`
	Total_count     int64           `db:"total_count"`
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
