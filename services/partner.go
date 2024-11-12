package services

import (
	"github.com/gin-gonic/gin"
	"github.com/sathirak/garm/internal/errx"
	"github.com/sathirak/garm/internal/jwt"
	"github.com/sathirak/garm/internal/validator"
	"github.com/sathirak/garm/models"
	"github.com/sathirak/garm/models/dto"
	"github.com/sathirak/garm/repository"
)

func SignUpUser(signUpDto *dto.SignUpUser) (*models.User, errx.Errx) {

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

	hash, salt, err := GenerateHashSalt(signUpDto.Password)
	if err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	userMeta, err := CreateUser(&dto.UserInit{
		FirstName:   signUpDto.FirstName,
		LastName:    signUpDto.LastName,
		Email:       signUpDto.Email,
		Locale:      signUpDto.Locale,
		ContactNo:   signUpDto.ContactNo,
		CountryCode: signUpDto.CountryCode,
	}, salt, hash)

	if err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	return userMeta, errx.Nil()
}

func SignInUser(signInDto *dto.SignInUser) (*models.User, errx.Errx) {

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

	userCredentials, err := repository.GetUserCredentials(signInDto.Email)

	if err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	if err := ValidateCredentials(userCredentials.Hash, userCredentials.Salt, signInDto.Password); !err.IsNil() {
		if err.ApiError == errx.ErrInvalidCredentials {
			if err := UpdateRetries(userCredentials.Retries, userCredentials.UserID); !err.IsNil() {
				return nil, err
			}
		}
		return nil, err
	}

	userMeta, err := repository.GetUserMeta(userCredentials.UserID)

	if err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	return userMeta, errx.Nil()
}

func ResetPasswordUser(resetDto *dto.ResetPasswordUser, c *gin.Context) errx.Errx {

	if _, err := jwt.Get(c); !err.IsNil() {
		return err
	}

	if err := validator.ValidatePassword(resetDto.NewPassword); err != nil {
		return errx.NewError(err, errx.ErrPasswordInvalid)
	}

	credentails, err := repository.GetUserCredentials(resetDto.Email)

	if err != nil {
		return errx.NewError(err, errx.ErrInternalServerErr)
	}

	if err := ValidateCredentials(credentails.Hash, credentails.Salt, resetDto.OldPassword); !err.IsNil() {
		return err
	}

	credentails.Hash, credentails.Salt, err = GenerateHashSalt(resetDto.NewPassword)
	if err != nil {
		return errx.NewError(err, errx.ErrInternalServerErr)
	}

	if err = repository.UpdateEmailPassword(credentails); err != nil {
		return errx.NewError(err, errx.ErrInternalServerErr)
	}

	return errx.Nil()
}
