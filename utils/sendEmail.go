package utils

import (
	"net/smtp"
)


func SendEmail() error {

	err := smtp.SendMail("smtp.gmail.com:587", smtp.PlainAuth("", "testsenderstori@gmail.com", "Hola123.", "smtp.gmail.com"),"", []string{"davidasb.developer@gmail.com"}, []byte("dummy message"))

	if err != nil {
		return err
	}

	return nil
}