package services

import (
	"time"

	"github.com/sathirak/garm/models"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/repository"
)

func CreateUser(userInit *dto.UserInit, salt string, hash string) (*models.User, error) {

	user := dto.UserCreate{
		VerifiedEmail: false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		UserInit:      *userInit,
	}

	userMeta, err := repository.CreateUser(&user, salt, hash)

	if err != nil {
		return userMeta, err
	}

	return userMeta, nil
}
