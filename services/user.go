package services

import (
	"time"

	"github.com/hotelbear/garm/models"
	"github.com/hotelbear/garm/models/dto"
	"github.com/hotelbear/garm/repository"
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
