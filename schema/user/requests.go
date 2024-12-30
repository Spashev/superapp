package handler

type CreateUserReq struct {
	Password    string  `json:"password"`
	Email       string  `json:"email"`
	FirstName   *string `json:"first_name"`
	LastName    *string `json:"last_name"`
	MiddleName  *string `json:"middle_name"`
	DateOfBirth *string `json:"date_of_birth"`
	PhoneNumber *string `json:"phone_number"`
	Avatar      *string `json:"avatar"`
	IIN         *string `json:"iin"`
	Role        string  `json:"role"`
	IsSuperuser bool    `json:"is_superuser"`
	IsAdmin     bool    `json:"is_admin"`
	IsStaff     bool    `json:"is_staff"`
	IsActive    bool    `json:"is_active"`
}
