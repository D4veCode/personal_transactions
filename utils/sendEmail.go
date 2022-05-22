package utils

import (
	"fmt"
	"net/smtp"

	"github.com/D4vecode/personal_transactions/config"
)

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)



func SendEmail(data CalculatedData, email string) error {

	EMAIL_SERVER, _ := config.GetEnv("EMAIL_SERVER")
  EMAIL_USER, _ := config.GetEnv("EMAIL_USER")
  EMAIL_PASSWORD, _ := config.GetEnv("EMAIL_PASSWORD")
  EMAIL_HOST, _ := config.GetEnv("EMAIL_HOST")

	body, err := CreateEmailTemplate(data)
	body = "Subject: Personal Transactions\n" + MIME + "\r\n" + body
	if err != nil {
		return err
	}

	err = smtp.SendMail(EMAIL_SERVER, smtp.PlainAuth("", EMAIL_USER, EMAIL_PASSWORD, EMAIL_HOST),"", []string{email}, []byte(body))

	if err != nil {
		return err
	}

	fmt.Println("Email sent!")
	return nil
}