package repository

import (
	"github.com/hotelbear/garm/internal/db"
	"github.com/hotelbear/garm/internal/errx"
	"github.com/hotelbear/garm/models"
)

func CreateMailLog(mailLog *models.MailLogTable) errx.Errx {
	if err := db.Get().Omit("id").Create(mailLog).Error; err != nil {
		return errx.NewError(err, errx.ErrDatabase)
	}
	return errx.Nil()
}
