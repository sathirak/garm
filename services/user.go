package services

import (
	"time"

	"github.com/sathirak/garm/models"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/pkg/ksuid"
	"github.com/sathirak/garm/repository"
)

func CreateUser(dto *dto.UserInit) (*models.UserMeta, error) {

	user := models.UserMeta{
		ID:            ksuid.Gen().String(),
		VerifiedEmail: false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		User: models.User{
			Locale:    dto.Locale,
			FirstName: dto.FirstName,
			LastName:  dto.LastName,
			Email:     dto.Email},
	}

	err := repository.CreateUser(&user)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
