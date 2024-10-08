package dto

type UserInit struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Locale    string `json:"locale"`
}

type SignUpEmailPassword struct {
	UserInit
	Password string `json:"password"`
}

type SignInEmailPassword struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
