package dtos

import "errors"

type UserResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserRequest struct {
	UserID int `json:"user_id"`
}

type UserCreate struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Type  string `json:"type"`
}

func (u *UserRequest) Validate() error {
	if u.UserID == 0 {
		return errors.New("user_id is required")
	}
	return nil
}
