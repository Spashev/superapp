package schema

type RegisterReq struct {
	Email       string `json:"email" validate:"required,email"`
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	MiddleName  string `json:"middle_name"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	DateOfBirth string `json:"date_of_birth" validate:"required,date"`
	Password    string `json:"password" validate:"required,min=6"`
	IIN         string `json:"iin"`
}
