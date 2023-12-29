package utils

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"os"
	"strconv"
)

func SendMail(to string, message string, subject string) error {
	m := gomail.NewMessage()
	from := os.Getenv("MAIL_FROM")
	port, _ := strconv.Atoi(os.Getenv("MAIL_PORT"))
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)
	d := gomail.NewDialer(os.Getenv("MAIL_HOST"), port, from, os.Getenv("MAIL_PASSWORD"))

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: false}
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
