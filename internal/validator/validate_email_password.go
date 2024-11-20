package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/hotelbear/garm/internal/errx"
	"github.com/hotelbear/garm/internal/logger"
	"github.com/hotelbear/garm/models"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

var Validate *validator.Validate

func ValidateSignUp(signUpDto *models.SignUpUserReq) errx.Errx {
	log := logger.Get()
	Validate = validator.New(validator.WithRequiredStructEnabled())

	err := Validate.Struct(signUpDto)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Info(err)
			return errx.NewError(err, errx.ErrInternalServer)
		}

		// from here you can create your own error messages in whatever language you wish
		return errx.NewError(nil, errx.ErrUnprocessableContent)
	}
	return errx.Nil()
}

func ValidatePassword(password string) errx.Errx {

	const minEntropyBits = 60
	err := passwordvalidator.Validate(password, minEntropyBits)

	if err != nil {
		return errx.NewError(err, errx.ErrInvalidCredentials)
	}

	return errx.Nil()
}

func CheckPasswordEntropy(password string) float64 {
	return passwordvalidator.GetEntropy(password)
}

func ValidateSignIn(signInDto *models.SignInUserReq) errx.Errx {
	log := logger.Get()
	Validate = validator.New(validator.WithRequiredStructEnabled())

	err := Validate.Struct(signInDto)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Info(err)
			return errx.NewError(err, errx.ErrInternalServer)
		}

		// from here you can create your own error messages in whatever language you wish
		return errx.NewError(nil, errx.ErrUnprocessableContent)
	}
	return errx.Nil()
}
