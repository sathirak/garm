package repository

import (
	"errors"

	"github.com/hotelbear/garm/internal/db"
	"github.com/hotelbear/garm/internal/errx"
	"github.com/hotelbear/garm/models"
	"gorm.io/gorm"
)

func GetMailTemplate(templateId int) (*models.MailTemplateTable, errx.Errx) {
	var mailTemplate models.MailTemplateTable

	err := db.Get().First(&mailTemplate, "id = ?", templateId).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errx.NewError(err, errx.ErrDatabaseRecordNotFound)
	}

	if err != nil {
		return nil, errx.NewError(err, errx.ErrDatabase)
	}

	return &mailTemplate, errx.Nil()
}
