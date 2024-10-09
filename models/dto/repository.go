package dto

import "time"

type AuthCredentials struct {
	ID             int       `json:"id"`
	AuthUserID     string    `json:"auth_user_id"`
	AuthMethodID   int       `json:"auth_method_id"`
	AuthIdentifier string    `json:"auth_identifier"`
	AuthSecret     string    `json:"auth_secret"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
