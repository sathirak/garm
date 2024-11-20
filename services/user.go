package services

import (
	"github.com/gin-gonic/gin"
	"github.com/hotelbear/garm/internal/errx"
	"github.com/hotelbear/garm/internal/hashing"
	"github.com/hotelbear/garm/internal/jwt"
	"github.com/hotelbear/garm/internal/validator"
	"github.com/hotelbear/garm/models"
	"github.com/hotelbear/garm/repository"
	"github.com/mitchellh/mapstructure"
)

func SignUpUser(signUpDto *models.SignUpUserReq) (*models.UserRes, errx.Errx) {

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

	hash, salt, err := hashing.GenerateHashSalt(signUpDto.Password)
	if err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	user := &models.UserTable{}
	if err := mapstructure.Decode(signUpDto, user); err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	userCredential := &models.UserCredentialTable{
		UserID: user.ID,
		Salt:   salt,
		Hash:   hash,
	}

	user.Status = "active"

	userMeta, err := repository.CreateUser(user, userCredential)

	if err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	return userMeta, errx.Nil()
}

func SignInUser(signInDto *models.SignInUserReq) (*models.UserRes, errx.Errx) {

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

	userWithCredentials, err := repository.GetUserCredential(signInDto.Email)

	if err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	if err := hashing.ValidateCredentials(userWithCredentials.Credential.Hash, userWithCredentials.Credential.Salt, signInDto.Password); !err.IsNil() {
		if err.ApiError == errx.ErrInvalidCredentials {
			if err := UpdateRetries(userWithCredentials.Credential.Retries, userWithCredentials.ID); !err.IsNil() {
				return nil, err
			}
		}
		return nil, err
	}

	userRes := &models.UserRes{}
	if err := mapstructure.Decode(userWithCredentials, userRes); err != nil {
		return nil, errx.NewError(err, errx.ErrInternalServerErr)
	}

	return userRes, errx.Nil()
}

func ResetPasswordUser(resetDto *models.ResetPasswordUserReq, c *gin.Context) errx.Errx {

	if _, err := jwt.Get(c); !err.IsNil() {
		return err
	}

	if err := validator.ValidatePassword(resetDto.NewPassword); err != nil {
		return errx.NewError(err, errx.ErrPasswordInvalid)
	}

	userWithCredential, err := repository.GetUserCredential(resetDto.Email)

	if err != nil {
		return errx.NewError(err, errx.ErrInternalServerErr)
	}

	if err := hashing.ValidateCredentials(userWithCredential.Credential.Hash, userWithCredential.Credential.Salt, resetDto.OldPassword); !err.IsNil() {
		return err
	}

	userWithCredential.Credential.Hash, userWithCredential.Credential.Salt, err = hashing.GenerateHashSalt(resetDto.NewPassword)
	if err != nil {
		return errx.NewError(err, errx.ErrInternalServerErr)
	}

	if err = repository.UpdateEmailPassword(&userWithCredential.Credential); err != nil {
		return errx.NewError(err, errx.ErrInternalServerErr)
	}

	return errx.Nil()
}
