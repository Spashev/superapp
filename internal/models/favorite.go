package models

type FavoriteProducts struct {
	Id        int `json:"id"`
	LikeId    int `json:"like_id"`
	ProductId int `json:"product_id"`
	UserId    int `json:"user_id"`
}

type Like struct {
	Id        int `json:"id"`
	LikeId    int `json:"like_id"`
	ProductId int `json:"product_id"`
	Count     int `json:"count"`
}
