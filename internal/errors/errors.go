package errors

import "errors"

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
	ErrInvalidUserData = errors.New("invalid user data")
)
