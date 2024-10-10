package services

import (
	"github.com/sathirak/garm/models"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/pkg/logger"
	"github.com/sathirak/garm/repository"

	"github.com/sathirak/garm/services/recipes"
)

func SignUpEmailPassword(signUpDto *dto.SignUpEmailPassword) (*models.UserMeta, error) {

	userDto := &dto.UserInit{
		FirstName: signUpDto.FirstName,
		LastName:  signUpDto.LastName,
		Email:     signUpDto.Email,
		Locale:    signUpDto.Locale,
	}

	user, err := CreateUser(userDto)

	if err != nil {
		return nil, err
	}

	if err = recipes.CreateEmailPassword(user.ID, signUpDto.Password); err != nil {
		return nil, err
	}

	return user, nil
}

func SignInEmailPassword(signInDto *dto.SignInEmailPassword) (*models.UserMeta, error) {

	credentails, err := repository.GetCredentialsEmailPassword(signInDto.Email)

	if err != nil {
		return nil, err
	}

	isValid := recipes.ValidateEmailPassword(credentails.AuthSecret, credentails.AuthIdentifier, signInDto.Password)

	logger.Get().Info("is valid: ", isValid)
	if !isValid {
		return nil, nil
	}

	user, err := repository.GetUserMeta(credentails.AuthUserID)

	if err != nil {
		return nil, err
	}

	return user, nil
}
