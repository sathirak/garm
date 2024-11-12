package models

// full user scope
type User struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	ContactNo     string `json:"contact_no"`
	CountryCode   string `json:"country_code"`
	Email         string `json:"email"`
	Locale        string `json:"locale"`
	ID            string `json:"id"`
	VerifiedEmail bool   `json:"verified_email"`
}

type UserAuthenticate struct {
	ID string `json:"id"`
}

type UserCredential struct {
	UserID  string
	Salt    string
	Hash    string
	Retries int
}

type PasswordCheck struct {
	Strength int `json:"strength"`
}
