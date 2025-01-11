package models

type ProductImages struct {
	ID        int    `json:"id"`
	Thumbnail string `json:"thumbnail"`
	Original  string `json:"original"`
	MimeType  string `json:"mimetype"`
	IsLabel   bool   `json:"is_label"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
}
