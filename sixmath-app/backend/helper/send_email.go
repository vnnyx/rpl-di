package helper

import (
	"bytes"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

func SendEmailTo(email string, buff *bytes.Buffer) error {
	m := gomail.NewMessage()
	m.SetHeader("From", "no-reply-sixmath@email.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Reset Password")
	m.SetBody("text/html", buff.String())
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	d := gomail.NewDialer(os.Getenv("MAIL_HOST"), port, os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"))
	err := d.DialAndSend(m)
	return err
}
