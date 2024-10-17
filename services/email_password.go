package services

import (
	"github.com/sathirak/garm/internal/errx"
	"github.com/sathirak/garm/internal/validator"
	"github.com/sathirak/garm/models"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/repository"

	"github.com/sathirak/garm/services/recipes"
)

func SignUpEmailPassword(signUpDto *dto.SignUpEmailPassword) (*models.UserMeta, errx.Errx) {

	IsAvailable, err := repository.IsEmailAvailable(signUpDto.Email)
	if err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	if !IsAvailable {
		return nil, errx.NewError(nil, errx.ErrEmailUnavailable)
	}

	if err = validator.ValidateSignUp(signUpDto); err != nil {
		return nil, errx.NewError(err, errx.ErrUnprocessableContent)
	}

	if err := validator.ValidatePassword(signUpDto.Password); err != nil {
		return nil, errx.NewError(err, errx.ErrPasswordInvalid)
	}

	user, err := CreateUser(&dto.UserInit{
		FirstName:   signUpDto.FirstName,
		LastName:    signUpDto.LastName,
		Email:       signUpDto.Email,
		Locale:      signUpDto.Locale,
		ContactNo:   signUpDto.ContactNo,
		CountryCode: signUpDto.CountryCode,
	})

	if err != nil {
		return nil, errx.NewError(nil, err)
	}

	hash, salt, err := recipes.CreateEmailPassword(user.ID, signUpDto.Password)
	if err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	if err = repository.CreateEmailPassword(user.ID, salt, hash); err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	return user, errx.NewError(nil, nil)
}

func SignInEmailPassword(signInDto *dto.SignInEmailPassword) (*models.UserMeta, errx.Errx) {

	if err := validator.ValidateSignIn(signInDto); err != nil {
		return nil, errx.NewError(err, errx.ErrUnprocessableContent)
	}

	if _, err := repository.IsEmailAvailable(signInDto.Email); err != nil {
		return nil, errx.NewError(err, errx.ErrEmailUnavailable)
	}

	credentails, err := repository.GetCredentialsEmailPassword(signInDto.Email)

	if err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	if isValid, err := recipes.ValidateEmailPassword(credentails.Hash, credentails.Salt, signInDto.Password); err != nil || !isValid {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	userMeta, err := repository.GetUserMeta(credentails.UserID)

	return userMeta, errx.NewError(err, errx.ErrInternalServerErr)
}
