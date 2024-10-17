package services

import (
	"github.com/sathirak/garm/internal/errors"
	"github.com/sathirak/garm/internal/validator"
	"github.com/sathirak/garm/models"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/repository"

	"github.com/sathirak/garm/services/recipes"
)

func SignUpEmailPassword(signUpDto *dto.SignUpEmailPassword) (*models.UserMeta, error) {

	if !validator.ValidateSignUp(signUpDto) {
		return nil, errors.ErrInvalidUserData
	}

	if err := validator.ValidatePassword(signUpDto.Password); err != nil {
		return nil, err
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
		return nil, err
	}

	hash, salt, err := recipes.CreateEmailPassword(user.ID, signUpDto.Password)
	if err != nil {
		return nil, err
	}

	if err = repository.CreateEmailPassword(user.ID, salt, hash); err != nil {
		return nil, err
	}

	return user, nil
}

func SignInEmailPassword(signInDto *dto.SignInEmailPassword) (*models.UserMeta, error) {

	if !validator.ValidateSignIn(signInDto) {
		return nil, errors.ErrInvalidUserData
	}

	if repository.IsEmailAvailable(signInDto.Email) {
		return nil, errors.ErrInvalidUserData
	}

	credentails, err := repository.GetCredentialsEmailPassword(signInDto.Email)

	if err != nil {
		return nil, err
	}

	if isValid, err := recipes.ValidateEmailPassword(credentails.Hash, credentails.Salt, signInDto.Password); err != nil || !isValid {
		return nil, err
	}

	return repository.GetUserMeta(credentails.UserID)
}
