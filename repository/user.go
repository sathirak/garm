package repository

import (
	"errors"

	"github.com/hotelbear/garm/internal/db"
	"github.com/hotelbear/garm/internal/errx"
	"github.com/hotelbear/garm/models"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func IsEmailAvailable(email string) (bool, errx.Errx) {
	err := db.Get().First(&models.UserTable{}, "email = ?", email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true, errx.Nil()
	} else if err != nil {
		return false, errx.NewError(err, errx.ErrDatabase)
	}

	return false, errx.Nil()
}

func CreateUser(user *models.UserTable, userCredential *models.UserCredentialTable) (*models.UserRes, errx.Errx) {
	conn := db.Get()

	err := conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("id", "is_email_verified", "status", "is_deleted", "created_at", "updated_at").Clauses(clause.Returning{}).Create(&user).Error; err != nil {
			return err
		}

		if err := tx.Create(userCredential).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, errx.NewError(err, errx.ErrDatabase)
	}

	userRes := &models.UserRes{}
	if err := mapstructure.Decode(user, userRes); err != nil {
		return nil, errx.NewError(err, errx.ErrParsing)
	}
	return userRes, errx.Nil()
}
