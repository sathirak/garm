package errx

import "errors"

type Errx struct {
	ApiError ApiErrx
	SvcError error
}

func NewError(svcError error, apiError ApiErrx) Errx {
	return Errx{
		ApiError: apiError,
		SvcError: svcError,
	}
}

func Nil() Errx {
	return Errx{
		ApiError: ApiErrx{nil, 200},
		SvcError: nil,
	}
}

func (e Errx) Error() string {
	return errors.Join(e.ApiError.Err, e.SvcError).Error()
}

func (e Errx) IsNil() bool {
	return e.ApiError.Err == nil && e.SvcError == nil
}

func (e Errx) GetApiError() error {
	return e.ApiError.Err
}

func (e Errx) GetSvcError() error {
	return e.SvcError
}
