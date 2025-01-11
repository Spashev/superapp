package models

type User struct {
	Email         string `json:"email"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Middle_name   string `json:"middle_name"`
	Date_of_birth string `json:"date_of_birth"`
	Phone_number  string `json:"phone_number"`
	Avatar        string `json:"avatar"`
	IIN           string `json:"iin"`
	Role          string `json:"role"`
	Is_active     string `json:"is_active"`
}

type OwnerProduct struct {
	Id           int64  `json:"id"`
	Email        string `json:"email"`
	First_name   string `json:"first_name"`
	Last_name    string `json:"last_name"`
	Middle_name  string `json:"middle_name"`
	Phone_number string `json:"phone_number"`
	Avatar       string `json:"avatar"`
}
