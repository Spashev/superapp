package models

// Category represents a country
// swagger:model
type Country struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// Category represents a city
// swagger:model
type City struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Postall_code string `json:"postall_code"`
}
