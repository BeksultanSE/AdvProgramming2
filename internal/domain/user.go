package domain

import (
	"errors"
	"regexp"
)

type User struct {
	ID    uint64
	Name  string
	Email string
}

func (u *User) Validate() error {
	if err := u.ValidateName(); err != nil {
		return err
	}

	if err := u.ValidateEmail(); err != nil {
		return err
	}

	return nil
}

func (u *User) ValidateName() error {
	if u.Name == "" {
		return errors.New("name cannot be empty")
	}
	if len(u.Name) < 3 {
		return errors.New("name must be at least 3 characters long")
	}
	return nil
}

func (u *User) ValidateEmail() error {
	if u.Email == "" {
		return errors.New("email cannot be empty")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(u.Email) {
		return errors.New("invalid email format")
	}

	return nil
}
