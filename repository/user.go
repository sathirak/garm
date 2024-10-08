package repository

import (
	"github.com/sathirak/garm/internal/db"
	"github.com/sathirak/garm/models"
	"github.com/sathirak/garm/models/dto"
)

func CreateUser(user *models.UserCreate) error {
	conn := db.Get()

	_, err := conn.Query("INSERT INTO auth_users (id, first_name, last_name, email, verified_email, locale, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);", user.ID, user.FirstName, user.LastName, user.Email, user.VerifiedEmail, user.Locale, user.CreatedAt, user.UpdatedAt)

	return err

}

func CreatePasswordMethod(userID string, salt string, hash string) error {
	conn := db.Get()

	_, err := conn.Query("INSERT INTO auth_credentials (auth_user_id, auth_method_id, auth_identifier, auth_secret) VALUES ($1, $2, $3, $4);", userID, 1, salt, hash)

	return err
}

func GetCredentialsEmailPassword(email string) (*dto.AuthCredentials, error) {
	conn := db.Get()

	var credentials dto.AuthCredentials

	err := conn.QueryRow(`
		SELECT ac.id, ac.auth_user_id, ac.auth_method_id, ac.auth_identifier, ac.auth_secret, ac.created_at, ac.updated_at
		FROM auth_credentials ac
		JOIN auth_users au ON ac.auth_user_id = au.id
		WHERE au.email = $1;`, email).Scan(
		&credentials.ID,
		&credentials.AuthUserID,
		&credentials.AuthMethodID,
		&credentials.AuthIdentifier,
		&credentials.AuthSecret,
		&credentials.CreatedAt,
		&credentials.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &credentials, nil
}

func GetUser(id string) (*models.User, error) {
	conn := db.Get()

	var user models.User

	err := conn.QueryRow(`
		SELECT first_name, last_name, email, locale FROM auth_users WHERE id = $1;`, id).Scan(
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Locale,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
