package utils

import (
	"fmt"
	"net/smtp"
)

const (
	MIME = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

func SendEmail(data CalculatedData, email string) error {
	body, err := CreateEmailTemplate(data)
	body = "Subject: Personal Transactions\n" + MIME + "\r\n" + body
	if err != nil {
		return err
	}

	err = smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", "testsenderstori@gmail.com", "Hola123.", "smtp.gmail.com"),"", []string{email}, []byte(body))

	if err != nil {
		return err
	}

	fmt.Println("Email sent!")
	return nil
}