package models

type SignUpUserReq struct {
	FirstName   string `json:"first_name" validate:"required,min=1,max=50,alpha"`
	LastName    string `json:"last_name" validate:"required,min=1,max=50,alpha"`
	ContactNo   string `json:"contact_no" validate:"required,min=1,max=15"`
	CountryCode string `json:"country_code" validate:"required,min=1,max=15"`
	Email       string `json:"email" validate:"required,email"`
	Locale      string `json:"locale" validate:"required,len=2,alpha"`
	Password    string `json:"password" validate:"required,min=8,max=100"`
}

type SignInUserReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=100"`
}

type ResetPasswordUserReq struct {
	Email       string `json:"email" validate:"required,email"`
	OldPassword string `json:"old_password" validate:"required,min=8,max=100"`
	NewPassword string `json:"new_password" validate:"required,min=8,max=100"`
}

type PasswordCheckReq struct {
	Password string `json:"password"`
}
