package validator

import (
	"github.com/go-playground/validator/v10"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/pkg/logger"
	passwordvalidator "github.com/wagslane/go-password-validator"
)

var validate *validator.Validate

func ValidateSignUp(signUpDto *dto.SignUpEmailPassword) bool {
	log := logger.Get()
	validate = validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(signUpDto)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Info(err)
			return false
		}

		for _, err := range err.(validator.ValidationErrors) {
			log.Info(err.Namespace())
			log.Info(err.Field())
			log.Info(err.StructNamespace())
			log.Info(err.StructField())
			log.Info(err.Tag())
			log.Info(err.ActualTag())
			log.Info(err.Kind())
			log.Info(err.Type())
			log.Info(err.Value())
			log.Info(err.Param())
			log.Info()
		}

		// from here you can create your own error messages in whatever language you wish
		return false
	}
	return true
}

func ValidatePassword(password string) error {

	const minEntropyBits = 60
	return passwordvalidator.Validate(password, minEntropyBits)
}

func ValidateSignIn(signInDto *dto.SignInEmailPassword) bool {
	log := logger.Get()
	validate = validator.New(validator.WithRequiredStructEnabled())

	err := validate.Struct(signInDto)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Info(err)
			return false
		}

		for _, err := range err.(validator.ValidationErrors) {
			log.Info(err.Namespace())
			log.Info(err.Field())
			log.Info(err.StructNamespace())
			log.Info(err.StructField())
			log.Info(err.Tag())
			log.Info(err.ActualTag())
			log.Info(err.Kind())
			log.Info(err.Type())
			log.Info(err.Value())
			log.Info(err.Param())
			log.Info()
		}

		// from here you can create your own error messages in whatever language you wish
		return false
	}
	return true
}
