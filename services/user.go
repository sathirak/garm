package services

import (
	"time"

	"github.com/sathirak/garm/models"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/pkg/ksuid"
	"github.com/sathirak/garm/repository"
	"github.com/sathirak/garm/services/methods"
)

func CreateUser(user *models.User, dto dto.User) error {

	*user = models.User{
		ID:            ksuid.Gen().String(),
		FirstName:     dto.FirstName,
		LastName:      dto.LastName,
		Email:         dto.Email,
		VerifiedEmail: false,
		Locale:        dto.Locale,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	err := repository.CreateUser(user)

	if err != nil {
		return err
	}

	return methods.CreatePassword(user.ID, "password")
}
