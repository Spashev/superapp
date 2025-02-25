package models

// Category represents a product images
// swagger:model
type ProductImages struct {
	ID        int    `json:"id"`
	Thumbnail string `json:"thumbnail"`
	Original  string `json:"original"`
	Mimetype  string `json:"mimetype"`
	Is_label  bool   `json:"is_label"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}
