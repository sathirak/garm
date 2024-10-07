package services

import (
	"time"

	"github.com/sathirak/garm/models"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/pkg/id"
)

func CreateUser(user *models.User, dto dto.User) error {

	*user = models.User{
		ID:            id.Gen().String(),
		FirstName:     dto.FirstName,
		LastName:      dto.LastName,
		Email:         dto.Email,
		VerifiedEmail: false,
		Locale:        dto.Locale,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	return nil
}
