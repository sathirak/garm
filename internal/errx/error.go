package errx

import "errors"

type Errx struct {
	ApiError error
	SvcError error
}

func NewError(svcError, appError error) Errx {
	return Errx{
		ApiError: appError,
		SvcError: svcError,
	}
}

func (e Errx) Error() string {
	return errors.Join(e.ApiError, e.SvcError).Error()
}

func (e Errx) IsNil() bool {
	return e.ApiError == nil && e.SvcError == nil
}

func (e Errx) GetApiError() error {
	return e.ApiError
}

func (e Errx) GetSvcError() error {
	return e.SvcError
}
