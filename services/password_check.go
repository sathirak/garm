package services

import (
	"math"

	"github.com/hotelbear/garm/internal/validator"
	"github.com/hotelbear/garm/models"
)

func PasswordCheck(password string) models.PasswordCheck {

	entropy := validator.CheckPasswordEntropy(password)
	return models.PasswordCheck{
		Strength: int(math.Floor(entropy)),
	}
}
