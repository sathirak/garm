package services

import (
	"github.com/sathirak/garm/internal/errx"
	"github.com/sathirak/garm/repository"
)

func UpdateRetries(retries int, userId string) errx.Errx {

	if retries >= 0 && retries < 5 {
		retries++

		if err := repository.UpdateRetries(retries, userId); err != nil {
			return errx.NewError(err, errx.ErrInternalServerErr)
		}
		return errx.Nil()
	}
	return errx.NewError(nil, errx.ErrInternalServerErr)
}
