package shared

import (
	"errors"
	"net/mail"
	"strings"
	"unicode/utf8"
)

func ValidateFirstName(firstName string) error {
	if strings.TrimSpace(firstName) == "" {
		return errors.New("first name is required")
	}
	return nil
}

func ValidateLastName(lastName string) error {
	if strings.TrimSpace(lastName) == "" {
		return errors.New("last name is required")
	}
	return nil
}

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

func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("Password is required...")
	}
	//unicode charecter count  =>utf8
	if utf8.RuneCountInString(password) < 8 {
		return errors.New("Password must be at least 8 charecters")
	}
	return nil
}
