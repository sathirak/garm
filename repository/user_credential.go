package repository

import (
	"errors"

	"github.com/hotelbear/garm/internal/db"
	"github.com/hotelbear/garm/internal/errx"
	"github.com/hotelbear/garm/models"
	"gorm.io/gorm"
)

func UpdateEmailPassword(userCredentials *models.UserCredentialTable) errx.Errx {
	err := db.Get().Model(userCredentials).Updates(models.UserCredentialTable{Hash: userCredentials.Hash, Salt: userCredentials.Salt}).Error
	return errx.NewError(err, errx.ErrDatabase)
}

func GetUserCredential(email string) (*models.UserWithCredentials, errx.Errx) {
	var user models.UserWithCredentials

	err := db.Get().Preload("C").First(&user, "email = ?", email).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errx.NewError(err, errx.ErrDatabaseRecordNotFound)
	}

	if err != nil {
		return nil, errx.NewError(err, errx.ErrDatabase)
	}

	return &user, errx.Nil()
}

func UpdateRetries(retries int, userId string) errx.Errx {
	err := db.Get().Model(&models.UserCredentialTable{}).
		Where("user_id = ?", userId).
		Update("retries", retries).
		Error

	if err != nil {
		return errx.NewError(err, errx.ErrDatabase)
	}

	return errx.Nil()
}
