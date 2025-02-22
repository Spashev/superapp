package models

type ProductCommentUser struct {
	Id         int    `json:"id"`
	Email      string `json:"email"`
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Avatar     string `json:"avatar"`
}

type ProductComment struct {
	ID         int                `json:"id"`
	Content    string             `json:"content"`
	Rating     int                `json:"rating"`
	User       ProductCommentUser `json:"user"`
	Created_at string             `json:"created_at"`
}
