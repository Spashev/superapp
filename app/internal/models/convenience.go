package models

// Category represents a product convenience
// swagger:model
type Convenience struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Slug string `json:"slug"`
}
