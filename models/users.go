package models

import "errors"

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Type  string `json:"type"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	if u.Type == "" {
		return errors.New("type is required")
	}

	return nil
}
