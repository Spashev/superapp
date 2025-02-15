package models

type User struct {
	Id            int    `json:"id"`
	Email         string `json:"email"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Middle_name   string `json:"middle_name"`
	Date_of_birth string `json:"date_of_birth"`
	Phone_number  string `json:"phone_number"`
	Avatar        string `json:"avatar"`
	IIN           string `json:"iin"`
	Is_active     bool   `json:"is_active"`
	Date_joined   string `json:"date_joined"`
}

type CreateUser struct {
	Id            int    `json:"id"`
	Email         string `json:"email"`
	First_name    string `json:"first_name"`
	Last_name     string `json:"last_name"`
	Middle_name   string `json:"middle_name"`
	Date_of_birth string `json:"date_of_birth"`
	Phone_number  string `json:"phone_number"`
	Avatar        string `json:"avatar"`
	IIN           string `json:"iin"`
	Role          string `json:"role"`
	Is_active     bool   `json:"is_active"`
	Password      string `json:"password"`
	Date_joined   string `json:"date_joined"`
}

type OwnerProduct struct {
	Id           int    `json:"id"`
	Email        string `json:"email"`
	First_name   string `json:"first_name"`
	Last_name    string `json:"last_name"`
	Middle_name  string `json:"middle_name"`
	Phone_number string `json:"phone_number"`
	Avatar       string `json:"avatar"`
}
