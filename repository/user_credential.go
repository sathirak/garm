package repository

import (
	"github.com/hotelbear/garm/internal/db"
	"github.com/hotelbear/garm/models"
)

func UpdateEmailPassword(userCredentials *models.UserCredentialTable) error {

	err := db.Get().Model(userCredentials).Updates(models.UserCredentialTable{Hash: userCredentials.Hash, Salt: userCredentials.Salt}).Error

	return err
}

func GetUserCredential(email string) (*models.UserWithCredentials, error) {
	var user models.UserWithCredentials

	err := db.Get().Preload("Credential").First(&user, "email = ?", email).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateRetries(retries int, userId string) error {
	return db.Get().Model(&models.UserCredentialTable{}).
		Where("user_id = ?", userId).
		Update("retries", retries).
		Error
}
