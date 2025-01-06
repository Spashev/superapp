package handler

type UserRes struct {
	ID          int64   `json:"id"`
	Email       string  `json:"email"`
	FirstName   *string `json:"first_name,omitempty"`
	LastName    *string `json:"last_name,omitempty"`
	MiddleName  *string `json:"middle_name,omitempty"`
	DateOfBirth *string `json:"date_of_birth,omitempty"`
	PhoneNumber *string `json:"phone_number,omitempty"`
	Avatar      *string `json:"avatar,omitempty"`
	IIN         *string `json:"iin,omitempty"`
	Role        string  `json:"role"`
	IsSuperuser bool    `json:"is_superuser"`
	IsAdmin     bool    `json:"is_admin"`
	IsStaff     bool    `json:"is_staff"`
	IsActive    bool    `json:"is_active"`
	CreatedAt   string  `json:"created_at"`
	UpdatedAt   string  `json:"updated_at"`
	LastLogin   *string `json:"last_login,omitempty"`
	DateJoined  string  `json:"date_joined"`
}
