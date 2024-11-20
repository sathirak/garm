package services

import (
	"math"

	"github.com/hotelbear/garm/internal/errx"
	"github.com/hotelbear/garm/internal/validator"
	"github.com/hotelbear/garm/models"
	"github.com/hotelbear/garm/repository"
)

func CheckPassword(password string) models.CheckPasswordRes {

	entropy := validator.CheckPasswordEntropy(password)
	return models.CheckPasswordRes{
		Strength: int(math.Floor(entropy)),
	}
}

func UpdateRetries(retries int, userId string) errx.Errx {

	if retries >= 0 && retries < 5 {
		retries++

		if err := repository.UpdateRetries(retries, userId); err != nil {
			return errx.NewError(err, errx.ErrInternalServerErr)
		}
		return errx.Nil()
	}
	return errx.NewError(nil, errx.ErrInternalServerErr)
}
