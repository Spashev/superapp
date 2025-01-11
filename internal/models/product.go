package models

type Product struct {
	Id                   int64            `json:"id"`
	Slug                 string           `json:"slug"`
	Name                 string           `json:"name"`
	Price_per_night      int64            `json:"price_per_night"`
	Price_per_week       int64            `json:"price_per_week"`
	Price_per_month      int64            `json:"price_per_month"`
	Owner                OwnerProduct     `json:"owner"`
	Rooms_qty            int64            `json:"rooms_qty"`
	Guest_qty            int64            `json:"guest_qty"`
	Bed_qty              int64            `json:"bed_qty"`
	Bedroom_qty          int64            `json:"bedroom_qty"`
	Toilet_qty           int64            `json:"toilet_qty"`
	Bath_qty             int64            `json:"bath_qty"`
	Description          string           `json:"description"`
	Country              string           `json:"country"`
	City                 string           `json:"city"`
	District             string           `json:"district"`
	Address              string           `json:"address"`
	Conveniences         []Convenience    `json:"convenience"`
	Type                 ProductType      `json:"type"`
	Lat                  string           `json:"lat"`
	Lng                  string           `json:"lng"`
	Like_count           int64            `json:"like_count"`
	Average_likes_rating float64          `json:"rating"`
	Phone_number         string           `json:"phone_number"`
	Is_new               bool             `json:"is_new"`
	Best_product         bool             `json:"best_product"`
	Promotion            bool             `json:"promotion"`
	Comments             []ProductComment `json:"comments"`
	Bookings             []string         `json:"bookings"`
	Images               []ProductImages  `json:"images"`
	Type_id              int64            `json:"type_id"`
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
