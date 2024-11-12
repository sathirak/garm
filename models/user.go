package models

import (
	"time"
)

// profile scope
type User struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	ContactNo   string `json:"contact_no"`
	CountryCode string `json:"country_code"`
	Email       string `json:"email"`
	Locale      string `json:"locale"`
}

// full user scope
type UserMeta struct {
	User
	ID            string    `json:"id"`
	VerifiedEmail bool      `json:"verified_email"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type UserAuthenticate struct {
	ID string `json:"id"`
}

type UserCredentials struct {
	UserID  string
	Salt    string
	Hash    string
	Retries int
}

type PasswordCheck struct {
	Strength int `json:"strength"`
}
