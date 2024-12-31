package models

type ProductImages struct {
	ID          int    `json:"id"`
	Thumbnail   string `json:"thumbnail"`
	Original    string `json:"original"`
	MimeType    string `json:"mimetype"`
	IsLabel     bool   `json:"is_label"`
	ImageWidth  int    `json:"width"`
	ImageHeight int    `json:"height"`
}

type ProductImagePaginate struct {
	ID          int    `json:"id"`
	Thumbnail   string `json:"thumbnail"`
	MimeType    string `json:"mimetype"`
	IsLabel     bool   `json:"is_label"`
	ImageWidth  int    `json:"width"`
	ImageHeight int    `json:"height"`
}
