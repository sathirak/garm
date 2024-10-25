package repository

import (
	"github.com/sathirak/garm/internal/db"
	"github.com/sathirak/garm/models/dto"
)

func CreateEmailPassword(userID string, salt string, hash string) error {
  conn := db.Get()

  _, err := conn.Query("INSERT INTO email_credential (user_id, salt, hash) VALUES ($1, $2, $3);", userID, salt, hash)

  return err
}

func UpdateEmailPassword(userID string, salt string, hash string) error {
  conn := db.Get()

  _, err := conn.Query("UPDATE email_credential SET salt = $1, hash = $2 WHERE user_id = $3;", salt, hash, userID)

  return err
}

func GetCredentialsEmailPassword(email string) (*dto.EmailCredentials, error) {
	conn := db.Get()

	var credentials dto.EmailCredentials

	err := conn.QueryRow(`
		SELECT ec.id, ec.user_id, ec.salt, ec.hash, ec.created_at, ec.updated_at
		FROM email_credential ec
		JOIN "user" u ON ec.user_id = u.id
		WHERE u.email = $1;`, email).Scan(
		&credentials.ID,
		&credentials.UserID,
		&credentials.Salt,
		&credentials.Hash,
		&credentials.CreatedAt,
		&credentials.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &credentials, nil
}
