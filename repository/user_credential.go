package repository

import (
	"github.com/hotelbear/garm/internal/db"
	"github.com/hotelbear/garm/models"
)

func CreateEmailPassword(userID string, salt string, hash string) error {
	userCredential := &models.UserCredentialDB{
		UserID: userID,
		Salt:   salt,
		Hash:   hash,
	}
	err := db.Get().Create(&userCredential).Error

	return err
}

func UpdateEmailPassword(userCredentials *models.UserCredential) error {

	userCredentialDB := &models.UserCredentialDB{
		UserID: userCredentials.UserID,
		Salt:   userCredentials.Salt,
		Hash:   userCredentials.Hash,
	}

	err := db.Get().Model(userCredentialDB).Updates(models.UserCredentialDB{Hash: userCredentials.Hash, Salt: userCredentials.Salt}).Error

	return err
}

func GetUserCredentials(email string) (*models.UserCredential, error) {
	var userCredential models.UserCredentialDB

	err := db.Get().
		Joins("JOIN \"user\" ON user_credential.user_id = \"user\".id").
		Where("\"user\".email = ?", email).
		First(&userCredential).Error

	if err != nil {
		return nil, err
	}

	return &models.UserCredential{
		UserID:  userCredential.UserID,
		Salt:    userCredential.Salt,
		Hash:    userCredential.Hash,
		Retries: userCredential.Retries,
	}, nil
}

func UpdateRetries(retries int, userId string) error {
	return db.Get().Model(&models.UserCredentialDB{}).
		Where("user_id = ?", userId).
		Update("retries", retries).
		Error
}
