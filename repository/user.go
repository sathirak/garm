package repository

import (
	"database/sql"

	"github.com/sathirak/garm/internal/db"
	"github.com/sathirak/garm/models"
)

func IsEmailAvailable(email string) (bool, error) {
	// If email is not in table returns true
	conn := db.Get()

	var existingEmail string

	err := conn.QueryRow(
		`SELECT "email" FROM "user" WHERE "email" = $1;`,
		email).Scan(&existingEmail)

	if err == sql.ErrNoRows {
		return true, nil
	}

	if err != nil {
		return false, err
	}

	// If there's an error (other than no rows) or if an email is found, it's not available
	return false, nil
}

func IsIDAvailable(id string) (bool, error) {
	// If ID is not in table returns true
	conn := db.Get()

	var existingID string

	err := conn.QueryRow(
		`SELECT id FROM "user" WHERE id = $1;`,
		id).Scan(&existingID)

	if err == sql.ErrNoRows {
		return true, nil
	}

	if err != nil {
		return false, err
	}

	// If there's an error (other than no rows) or if an id is found, it's not available
	return false, nil
}

func CreateUser(user *models.UserMeta) error {
	conn := db.Get()

	row := conn.QueryRow(
		`INSERT INTO "user" (id, first_name, last_name, email, is_email_verified, locale, contact_no, country_code, created_at, updated_at)
     VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
     RETURNING id, first_name, last_name, email, is_email_verified, locale, contact_no, country_code, created_at, updated_at;`,
		user.ID, user.FirstName, user.LastName, user.Email, user.VerifiedEmail, user.Locale, user.ContactNo, user.CountryCode, user.CreatedAt, user.UpdatedAt)

	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.VerifiedEmail, &user.Locale, &user.ContactNo, &user.CountryCode, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return err
	}

	return nil
}

func GetUserMeta(id string) (*models.UserMeta, error) {
	conn := db.Get()

	var user models.UserMeta

	err := conn.QueryRow(`
		SELECT first_name, last_name, email, locale, id, is_email_verified, created_at, updated_at FROM "user" WHERE id = $1;`, id).Scan(
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
