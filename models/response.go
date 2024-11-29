package models

type UserRes struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	ContactNo     string `json:"contact_no"`
	CountryCode   string `json:"country_code"`
	Email         string `json:"email"`
	Locale        string `json:"locale"`
	ID            string `json:"id"`
	VerifiedEmail bool   `json:"is_email_verified"`
}

type UserAuthenticateRes struct {
  ID            string `json:"id,omitempty"`
  VerifiedEmail *bool  `json:"is_email_verified,omitempty"`
  FirstName     string `json:"first_name,omitempty"`
  LastName      string `json:"last_name,omitempty"`
  ContactNo     string `json:"contact_no,omitempty"`
  CountryCode   string `json:"country_code,omitempty"`
  Email         string `json:"email,omitempty"`
  Locale        string `json:"locale,omitempty"`
}

type UserCredentialRes struct {
	UserID  string
	Salt    string
	Hash    string
	Retries int
}

type CheckPasswordRes struct {
	Strength int `json:"strength"`
}
