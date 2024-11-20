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
	ErrInternalServer = ApiErrx{errors.New("internal server error"), http.StatusInternalServerError}

	// Database errors
	ErrDatabase               = ApiErrx{errors.New("internal server error"), http.StatusInternalServerError}
	ErrDatabaseRecordNotFound = ApiErrx{errors.New("content not found"), http.StatusNotFound}

	ErrUnauthenticated      = ApiErrx{errors.New("unauthenticated request"), http.StatusUnauthorized}
	ErrUnprocessableContent = ApiErrx{errors.New("unprocessable content"), http.StatusUnprocessableEntity}

	ErrInvalidCredentials = ApiErrx{errors.New("email or password doesn't match"), http.StatusBadRequest}

	ErrInvalidBearerHeader = ApiErrx{errors.New("invalid bearer header"), http.StatusBadRequest}
	ErrInvalidToken        = ApiErrx{errors.New("invalid jwt token"), http.StatusUnauthorized}

	ErrParsing                    = ApiErrx{errors.New("error processing request"), http.StatusInternalServerError}
	ErrEmailUnavailable           = ApiErrx{errors.New("email unavailable"), http.StatusConflict}
	ErrPasswordInvalid            = ApiErrx{errors.New("password invalid"), http.StatusBadRequest}
	ErrMissingOrMalformedApiToken = ApiErrx{errors.New("missing or malformed api token"), http.StatusBadRequest}
)
