package repository

import (
	"database/sql"

	"github.com/sathirak/garm/internal/db"
	"github.com/sathirak/garm/models"
)

func CheckUserAvailablityEmail(email string) bool {
	// If email is not in table returns true
	conn := db.Get()

	var existingEmail string

	err := conn.QueryRow(
		`SELECT email FROM auth_users WHERE email = $1;`,
		email).Scan(&existingEmail)

	if err == sql.ErrNoRows {
		return err == sql.ErrNoRows
	}

	// If there's an error (other than no rows) or if an email is found, it's not available
	return false
}

func CheckUserAvailablityIdEmail(id, email string) bool {
	conn := db.Get()

	var existingID string
	var existingEmail string

	err := conn.QueryRow(
		`SELECT id, email FROM auth_users WHERE id = $1 OR email = $2;`,
		id, email).Scan(&existingID, &existingEmail)

	if err != nil && err != sql.ErrNoRows {
		return false
	}

	if existingID != "" || existingEmail != "" {
		return false
	}
	return true
}

func CreateUser(user *models.UserMeta) error {
	conn := db.Get()

	row := conn.QueryRow(
		`INSERT INTO auth_users (id, first_name, last_name, email, verified_email, locale, created_at, updated_at)
         VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
         RETURNING id, first_name, last_name, email, verified_email, locale, created_at, updated_at;`,
		user.ID, user.FirstName, user.LastName, user.Email, user.VerifiedEmail, user.Locale, user.CreatedAt, user.UpdatedAt)

	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.VerifiedEmail, &user.Locale, &user.CreatedAt, &user.UpdatedAt)

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

func GetUserMeta(id string) (*models.UserMeta, error) {
	conn := db.Get()

	var user models.UserMeta

	err := conn.QueryRow(`
		SELECT first_name, last_name, email, locale, id, verified_email, created_at, updated_at FROM auth_users WHERE id = $1;`, id).Scan(
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Locale,
		&user.ID,
		&user.VerifiedEmail,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
