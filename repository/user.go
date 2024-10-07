package repository

import (
	"github.com/sathirak/garm/internal/db"
	"github.com/sathirak/garm/models"
)

func CreateUser(user *models.User) error {
	conn := db.Get()

	_, err := conn.Query("INSERT INTO auth_users (id, first_name, last_name, email, verified_email, locale, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", user.ID, user.FirstName, user.LastName, user.Email, user.VerifiedEmail, user.Locale, user.CreatedAt, user.UpdatedAt)

	return err

}

func CreatePasswordCredentails(userID string, userPasswordHash string) error {
	conn := db.Get()

	_, err := conn.Query("INSERT INTO auth_credentials (auth_user_id, auth_method_id, auth_secret) VALUES ($1, $2, $3)", userID, 1, userPasswordHash)

	return err
}