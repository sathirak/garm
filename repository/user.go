package repository

import (
	"github.com/sathirak/garm/internal/db"
	"github.com/sathirak/garm/models"
)

func CreateUser(user *models.UserCreate)  error {
	conn := db.Get()

	row := conn.QueryRow(
		`INSERT INTO auth_users (id, first_name, last_name, email, verified_email, locale, created_at, updated_at)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		 RETURNING id, first_name, last_name, email, verified_email, locale, created_at, updated_at;`,
		user.ID, user.FirstName, user.LastName, user.Email, user.VerifiedEmail, user.Locale, user.CreatedAt, user.UpdatedAt)

	err := row.Scan(user.ID, user.FirstName, user.LastName, user.Email, user.VerifiedEmail, user.Locale, user.CreatedAt, user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
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
