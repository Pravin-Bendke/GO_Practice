package shared

import (
	"errors"
	"net/mail"
	"strings"
)

func ValidateEmail(email string) error {
	email = strings.TrimSpace(email)

	if email == "" {
		return errors.New("Email is Required...")
	}
	parsedEmail, err := mail.ParseAddress(email)

	if err != nil || parsedEmail.Address != email {
		return errors.New("Invalid email format")
	}
	return nil
}

func validatePassword(password string) error {
	if password == "" {
		return errors.New("Password is required...")
	}
	return nil
}
