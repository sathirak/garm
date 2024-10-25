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

	if err := validator.ValidateSignUp(signUpDto); !err.IsNil() {
		return nil, err
	}

	IsAvailable, err := repository.IsEmailAvailable(signUpDto.Email)
	if err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	if !IsAvailable {
		return nil, errx.NewError(nil, errx.ErrEmailUnavailable)
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
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	hash, salt, err := recipes.GenerateHashSalt(signUpDto.Password)
	if err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	if err = repository.CreateEmailPassword(user.ID, salt, hash); err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	return user, errx.Nil()
}

func SignInEmailPassword(signInDto *dto.SignInEmailPassword) (*models.UserMeta, errx.Errx) {

	if err := validator.ValidateSignIn(signInDto); !err.IsNil() {
		return nil, err
	}
	IsAvailable, err := repository.IsEmailAvailable(signInDto.Email)
	if err != nil {
		return nil, errx.NewError(err, errx.ErrEmailUnavailable)
	}

	if IsAvailable {
		return nil, errx.NewError(nil, errx.ErrEmailUnavailable)
	}

	credentails, err := repository.GetCredentialsEmailPassword(signInDto.Email)

	if err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	if err := recipes.ValidateEmailPassword(credentails.Hash, credentails.Salt, signInDto.Password); !err.IsNil() {
		return nil, err
	}

	userMeta, err := repository.GetUserMeta(credentails.UserID)

	if err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	return userMeta, errx.Nil()
}

func ResetEmailPassword(resetDto *dto.ResetEmailCredentials, userID string) errx.Errx {

	if err := validator.ValidatePassword(resetDto.NewPassword); err != nil {
		return errx.NewError(err, errx.ErrPasswordInvalid)
	}

	IsAvailable, err := repository.IsIDAvailable(userID)
	if err != nil {
		return errx.NewError(err, errx.ErrEmailUnavailable)
	}

	if IsAvailable {
		return errx.NewError(nil, errx.ErrEmailUnavailable)
	}

	credentails, err := repository.GetCredentialsEmailPassword(resetDto.Email)

	if err != nil {
		return errx.NewError(err, errx.ErrInternalServerErr)
	}

	if credentails.UserID != userID {
		return errx.NewError(err, errx.ErrUnprocessableContent)
	}

	if err := recipes.ValidateEmailPassword(credentails.Hash, credentails.Salt, resetDto.OldPassword); !err.IsNil() {
		return err
	}

	hash, salt, err := recipes.GenerateHashSalt(resetDto.NewPassword)
	if err != nil {
		return errx.NewError(err, errx.ErrInternalServerErr)
	}

	if err = repository.UpdateEmailPassword(userID, salt, hash); err != nil {
		return errx.NewError(err, errx.ErrInternalServerErr)
	}

	return errx.Nil()
}
