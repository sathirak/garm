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
	ErrInternalServerErr    = ApiErrx{errors.New("internal server error"), http.StatusInternalServerError}
	ErrUnauthenticated      = ApiErrx{errors.New("unauthenticated request"), http.StatusUnauthorized}
	ErrUnprocessableContent = ApiErrx{errors.New("unprocessable content"), http.StatusUnprocessableEntity}

	ErrInvalidCredentials = ApiErrx{errors.New("email or password doesn't match"), http.StatusBadRequest}

	ErrInvalidBearerHeader = ApiErrx{errors.New("invalid bearer header"), http.StatusBadRequest}
	ErrInvalidToken        = ApiErrx{errors.New("invalid jwt token"), http.StatusUnauthorized}
	ErrInvalidTokenClaims  = ApiErrx{errors.New("invalid jwt token claims"), http.StatusUnauthorized}

	ErrEmailUnavailable           = ApiErrx{errors.New("email unavailable"), http.StatusConflict}
	ErrPasswordInvalid            = ApiErrx{errors.New("password invalid"), http.StatusBadRequest}
	ErrInvalidUserData            = ApiErrx{errors.New("invalid user data"), http.StatusBadRequest}
	ErrMissingOrMalformedApiToken = ApiErrx{errors.New("missing or malformed api token"), http.StatusBadRequest}
)
