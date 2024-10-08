package models

import (
	"time"
)

// Profile scope
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Locale    string `json:"locale"`
}

// Create scope
type UserCreate struct {
	User
	ID            string    `json:"id"`
	VerifiedEmail bool      `json:"verified_email"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
