package services

import (
	"math"

	"github.com/sathirak/garm/internal/validator"
	"github.com/sathirak/garm/models"
)

func PasswordCheck(password string) models.PasswordCheck {

	entropy := validator.CheckPasswordEntropy(password)
	return models.PasswordCheck{
		Strength: int(math.Floor(entropy)),
	}
}
