package dto

import "time"

type UserInit struct {
	FirstName   string `json:"first_name" validate:"required,min=1,max=50,alpha"`
	LastName    string `json:"last_name" validate:"required,min=1,max=50,alpha"`
	ContactNo   string `json:"contact_no" validate:"required,min=1,max=15"`
	CountryCode string `json:"country_code" validate:"required,min=1,max=15"`
	Email       string `json:"email" validate:"required,email"`
	Locale      string `json:"locale" validate:"required,len=2,alpha"`
}

type UserCreate struct {
	UserInit
	VerifiedEmail bool
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type SignUpUser struct {
	UserInit
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type SignInUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type ResetPasswordUser struct {
	Email       string `json:"email" validate:"required,email"`
	OldPassword string `json:"old_password" validate:"required,min=8,max=100"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=100"`
}

type PasswordCheck struct {
	Password string `json:"password"`
}
