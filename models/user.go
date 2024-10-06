package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID
	FirstName     string
	LastName      string
	Email         string
	VerifiedEmail bool
	Locale        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
