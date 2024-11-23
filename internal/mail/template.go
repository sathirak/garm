package mail

import (
	"bytes"
	html "html/template"
	text "text/template"

	"github.com/hotelbear/garm/internal/errx"
	"github.com/hotelbear/garm/repository"
)

func Template(templateId int, data interface{}) (string, string, string, errx.Errx) {
	mailTemplate, err := repository.GetMailTemplate(templateId)

	if !err.IsNil() {
		return "", "", "", err
	}

	subject, parseErr := text.New("subject").Parse(mailTemplate.Subject)
	if parseErr != nil {
		return "", "", "", errx.NewError(parseErr, errx.ErrInternalServer)
	}

	htmlTemplate, parseErr := html.New("html").Parse(mailTemplate.HTMLBody)
	if parseErr != nil {
		return "", "", "", errx.NewError(parseErr, errx.ErrInternalServer)
	}

	textTemplate, parseErr := text.New("text").Parse(mailTemplate.Body)
	if parseErr != nil {
		return "", "", "", errx.NewError(parseErr, errx.ErrInternalServer)
	}

	var htmlBuffer, textBuffer, subjectBuffer bytes.Buffer
	if err := htmlTemplate.Execute(&htmlBuffer, data); err != nil {
		return "", "", "", errx.NewError(err, errx.ErrInternalServer)
	}
	if err := textTemplate.Execute(&textBuffer, data); err != nil {
		return "", "", "", errx.NewError(err, errx.ErrInternalServer)
	}

	if err := subject.Execute(&subjectBuffer, data); err != nil {
		return "", "", "", errx.NewError(err, errx.ErrInternalServer)
	}

	return subjectBuffer.String(), textBuffer.String(), htmlBuffer.String(), errx.Nil()
}
