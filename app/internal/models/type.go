package models

// Category represents a product type
// swagger:model
type ProductType struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
}
