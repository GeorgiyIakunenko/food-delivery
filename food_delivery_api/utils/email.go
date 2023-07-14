package utils

import (
	"fmt"
	"food_delivery/config"
	"net/smtp"
)

func SendResetCodeEmail(email string, resetCode string, cfg *config.Config) error {
	from := cfg.Email
	password := cfg.EmailPassword

	to := email
	subject := "Password Reset Code"
	body := fmt.Sprintf("Your password reset code is: %s", resetCode)

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		body

	auth := smtp.PlainAuth("", from, password, "smtp.gmail.com")

	err := smtp.SendMail("smtp.gmail.com:587", auth, from, []string{to}, []byte(msg))
	if err != nil {
		return err
	}

	return nil
}
