package models

type ProductCommentUser struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}

type ProductComment struct {
	ID      int                `json:"id"`
	Content string             `json:"content"`
	Rating  int                `json:"rating"`
	User    ProductCommentUser `json:"user"`
}
