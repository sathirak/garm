package repository

import (
	"github.com/sathirak/garm/internal/db"
	"github.com/sathirak/garm/models/dto"
)

func CreateEmailPassword(userID string, salt string, hash string) error {
	conn := db.Get()

	_, err := conn.Query("INSERT INTO user_credential (user_id, salt, hash) VALUES ($1, $2, $3);", userID, salt, hash)

	return err
}

func UpdateEmailPassword(userID string, salt string, hash string) error {
	conn := db.Get()

	_, err := conn.Query("UPDATE user_credential SET salt = $1, hash = $2 WHERE user_id = $3;", salt, hash, userID)

	return err
}

func GetUserCredentials(email string) (*dto.UserCredentials, error) {
	conn := db.Get()

	var credentials dto.UserCredentials

	err := conn.QueryRow(`
		SELECT uc.user_id, uc.salt, uc.hash, uc.created_at, uc.updated_at
		FROM user_credential uc
		JOIN "user" u ON uc.user_id = u.id
		WHERE u.email = $1;`, email).Scan(
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
