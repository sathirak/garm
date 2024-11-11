package dto

import (
	"time"
)

type UserCredentials struct {
	UserID    string    `json:"user_id"`
	Salt      string    `json:"salt"`
	Hash      string    `json:"hash"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
