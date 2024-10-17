package errx

import (
	"errors"
)

var ErrInternalServerErr = errors.New("internal server error")
var ErrUnauthenticated = errors.New("unauthenticated request")
var ErrUnprocessableContent = errors.New("unprocessable content")
// package recipes
var ErrHashDoesntMatch = errors.New("hash doesn't match")

// package jwt
var (
	ErrInvalidBearerHeader = errors.New("invalid bearer header")
	ErrInvalidToken        = errors.New("invalid jwt token")
	ErrInvalidTokenClaims  = errors.New("invalid jwt token claims")
)

// package services email password
var (
  ErrEmailUnavailable            = errors.New("email unavailable")
  ErrPasswordInvalid            = errors.New("password invalid")
	ErrInvalidUserData            = errors.New("invalid user data")
	ErrMissingOrMalformedApiToken = errors.New("missing or malformed api token")
)
