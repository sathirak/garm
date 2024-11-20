package errx

import (
	"errors"
	"net/http"
)

type ApiErrx struct {
	Err        error
	StatusCode int
}

var (
	ErrInternalServer = ApiErrx{errors.New("unexpected error occurred"), http.StatusInternalServerError}

	// Database errors
	ErrDatabase               = ApiErrx{errors.New("database error"), http.StatusInternalServerError}
	ErrDatabaseRecordNotFound = ApiErrx{errors.New("resource not found"), http.StatusNotFound}

	// Authentication errors
	ErrUnauthenticated            = ApiErrx{errors.New("sign in required"), http.StatusUnauthorized}
	ErrEmailUnavailable           = ApiErrx{errors.New("email already in use"), http.StatusConflict}
	ErrMissingOrMalformedApiToken = ApiErrx{errors.New("invalid api token"), http.StatusBadRequest}
	ErrInvalidCredentials         = ApiErrx{errors.New("invalid email or password"), http.StatusBadRequest}
	ErrInvalidToken               = ApiErrx{errors.New("session expired"), http.StatusUnauthorized}

	// Validation errors
	ErrInvalidBearerHeader  = ApiErrx{errors.New("invalid auth header"), http.StatusBadRequest}
	ErrUnprocessableContent = ApiErrx{errors.New("invalid input"), http.StatusUnprocessableEntity}
)
