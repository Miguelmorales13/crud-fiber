package email

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	template2 "html/template"
	"net/smtp"
	"os"
)

var util *EmailUtil

type EmailUtil struct {
}

func Util() *EmailUtil {
	if util == nil {
		util = &EmailUtil{}
	}
	return util
}

func (u *EmailUtil) GenTemplate(render string, options interface{}) (template string, err error) {
	var t *template2.Template
	t, err = t.ParseFiles("utils/email/templates/" + render + ".email.html")
	if err != nil {
		return
	}
	buff := new(bytes.Buffer)
	err = t.Execute(buff, options)
	if err != nil {
		return
	}
	template = buff.String()
	return
}
func (u *EmailUtil) SendEmail(to []string, template string, subject string) (err error) {
	auth := smtp.PlainAuth("", os.Getenv("EMAIL_FROM"), os.Getenv("EMAIL_PASSWORD"), os.Getenv("EMAIL_HOST"))
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := "Subject: " + subject + "\n" + mime + "\n\n" + template
	err = smtp.SendMail(os.Getenv("EMAIL_HOST")+":"+os.Getenv("EMAIL_PORT"), auth, os.Getenv("EMAIL_FROM"), []string{"cacahuatisimo13@protonmail.com"}, []byte(msg))
	return
}

func (u *EmailUtil) SendRecoveryPassword(to []string, password string) (err error) {
	template, err := u.GenTemplate("recovery-password", fiber.Map{"Password": password})
	if err != nil {
		return
	}
	err = u.SendEmail(to, template, "Recovery password")
	return
}
