package mail

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/smtp"
	"strings"
	"time"

	"github.com/hotelbear/garm/internal/config"
	"github.com/hotelbear/garm/internal/errx"
	"github.com/hotelbear/garm/models"
	"github.com/hotelbear/garm/repository"
)

func SendMail(to []string, from string, templateId int, data interface{}) errx.Errx {

	cfg := config.Get()

	subject, textBody, htmlBody, templateErr := Template(templateId, data)

	if !templateErr.IsNil() {
		return templateErr
	}

	auth := smtp.PlainAuth("", cfg.SMTP.Username, cfg.SMTP.Password, cfg.SMTP.Host)

	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "From: %s\n", from)
	fmt.Fprintf(buf, "To: %s\n", strings.Join(to, ", "))
	fmt.Fprintf(buf, "Subject: %s\n", subject)
	fmt.Fprintf(buf, "MIME-Version: 1.0\n")

	writer := multipart.NewWriter(buf)
	fmt.Fprintf(buf, "Content-Type: multipart/alternative; boundary=%s\n\n", writer.Boundary())

	if textBody != "" {
		part, _ := writer.CreatePart(map[string][]string{
			"Content-Type": {"text/plain; charset=UTF-8"},
		})
		_, err := part.Write([]byte(textBody))

		if err != nil {
			return errx.NewError(templateErr, errx.ErrInternalServer)
		}
	}

	if htmlBody != "" {
		part, _ := writer.CreatePart(map[string][]string{
			"Content-Type": {"text/html; charset=UTF-8"},
		})
		_, err := part.Write([]byte(htmlBody))

		if err != nil {
			return errx.NewError(templateErr, errx.ErrInternalServer)
		}
	}

	writer.Close()

	err := smtp.SendMail(
		cfg.SMTP.Host+":"+cfg.SMTP.Port,
		auth,
		from,
		to,
		buf.Bytes(),
	)

	mailLog := &models.MailLogTable{
		SentAt:         time.Now(),
		RecepientEmail: to[0],
		Data:           fmt.Sprintf("%v", data),
		TemplateID:     templateId,
	}

	if err != nil {
		mailLog.Status = "failed"
		if err := repository.CreateMailLog(mailLog); !err.IsNil() {
			return err
		}
		return errx.NewError(templateErr, errx.ErrInternalServer)
	}

	mailLog.Status = "sent"
	if err := repository.CreateMailLog(mailLog); !err.IsNil() {
		return err
	}
	return errx.Nil()
}
