package repository

import (
	"errors"

	"github.com/hotelbear/garm/internal/db"
	"github.com/hotelbear/garm/models"
	"github.com/hotelbear/garm/models/dto"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func IsEmailAvailable(email string) (bool, error) {
	err := db.Get().First(&models.UserDB{}, "email = ?", email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true, nil
	} else if err != nil {
		return false, err
	}

	return false, nil
}

func IsIDAvailable(id string) (bool, error) {
	err := db.Get().First(&models.UserDB{}, "id = ?", id).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return true, nil
	} else if err != nil {
		return false, err
	}

	return false, nil
}

func GetUserMeta(id string) (*models.User, error) {
	var user models.UserDB
	if err := db.Get().First(&user, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &models.User{
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

func CreateUser(user *dto.UserCreate, salt string, hash string) (*models.User, error) {
	conn := db.Get()

	userGorm := models.UserDB{
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		Email:         user.Email,
		Locale:        user.Locale,
		ContactNo:     user.ContactNo,
		CountryCode:   user.CountryCode,
		VerifiedEmail: user.VerifiedEmail,
	}

	err := conn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Omit("id", "is_email_verified", "status", "is_deleted", "created_at", "updated_at").Clauses(clause.Returning{}).Create(&userGorm).Error; err != nil {
			return err
		}

		userCredential := &models.UserCredentialDB{
			UserID: userGorm.ID,
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

	return &models.User{
		FirstName:     userGorm.FirstName,
		LastName:      userGorm.LastName,
		Email:         userGorm.Email,
		Locale:        userGorm.Locale,
		ContactNo:     userGorm.ContactNo,
		CountryCode:   userGorm.CountryCode,
		ID:            userGorm.ID,
		VerifiedEmail: userGorm.VerifiedEmail,
	}, nil
}
