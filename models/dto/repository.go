package dto

import (
	"time"

	"github.com/google/uuid"
)

type EmailCredentials struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	Salt      string    `json:"salt"`
	Hash      string    `json:"hash"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
