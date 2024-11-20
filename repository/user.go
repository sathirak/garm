package repository

import (
	"errors"

	"github.com/hotelbear/garm/internal/db"
	"github.com/hotelbear/garm/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func IsEmailAvailable(email string) (bool, error) {
	err := db.Get().First(&models.UserTable{}, "email = ?", email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true, nil
	} else if err != nil {
		return false, err
	}

	return false, nil
}

func IsIDAvailable(id string) (bool, error) {
	err := db.Get().First(&models.UserTable{}, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true, nil
	} else if err != nil {
		return false, err
	}

	return false, nil
}

func GetUserMeta(id string) (*models.UserRes, error) {
	var user models.UserTable
	if err := db.Get().First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &models.UserRes{
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		Locale:        user.Locale,
		ContactNo:     user.ContactNo,
		CountryCode:   user.CountryCode,
		ID:            user.ID,
		VerifiedEmail: user.VerifiedEmail,
	}, nil
}

func CreateUser(user *models.UserTable, salt string, hash string) (*models.UserRes, error) {
	conn := db.Get()

	err := conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("id", "is_email_verified", "status", "is_deleted", "created_at", "updated_at").Clauses(clause.Returning{}).Create(&user).Error; err != nil {
			return err
		}

		userCredential := &models.UserCredentialTable{
			UserID: user.ID,
			Salt:   salt,
			Hash:   hash,
		}

		if err := tx.Create(userCredential).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &models.UserRes{
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		Locale:        user.Locale,
		ContactNo:     user.ContactNo,
		CountryCode:   user.CountryCode,
		ID:            user.ID,
		VerifiedEmail: user.VerifiedEmail,
	}, nil
}
