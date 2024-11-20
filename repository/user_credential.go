package repository

import (
	"github.com/hotelbear/garm/internal/db"
	"github.com/hotelbear/garm/models"
)

func CreateEmailPassword(userID string, salt string, hash string) error {
	userCredential := &models.UserCredentialTable{
		UserID: userID,
		Salt:   salt,
		Hash:   hash,
	}
	err := db.Get().Create(&userCredential).Error

	return err
}

func UpdateEmailPassword(userCredentials *models.UserCredentialRes) error {

	userCredentialDB := &models.UserCredentialTable{
		UserID: userCredentials.UserID,
		Salt:   userCredentials.Salt,
		Hash:   userCredentials.Hash,
	}

	err := db.Get().Model(userCredentialDB).Updates(models.UserCredentialTable{Hash: userCredentials.Hash, Salt: userCredentials.Salt}).Error

	return err
}

func GetUserCredentials(email string) (*models.UserCredentialRes, error) {
	var userCredential models.UserCredentialTable

	err := db.Get().
		Joins("JOIN \"user\" ON user_credential.user_id = \"user\".id").
		Where("\"user\".email = ?", email).
		First(&userCredential).Error

	if err != nil {
		return nil, err
	}

	return &models.UserCredentialRes{
		UserID:  userCredential.UserID,
		Salt:    userCredential.Salt,
		Hash:    userCredential.Hash,
		Retries: userCredential.Retries,
	}, nil
}

func UpdateRetries(retries int, userId string) error {
	return db.Get().Model(&models.UserCredentialTable{}).
		Where("user_id = ?", userId).
		Update("retries", retries).
		Error
}
