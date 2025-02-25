package models

// Category represents a product category
type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Slug string `json:"slug"`
}
