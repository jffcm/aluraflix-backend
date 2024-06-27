package handler

import "errors"

type CreateUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *CreateUserRequest) Validate() error {
	if r.Username == "" {
		return errors.New("username is required")
	}

	if r.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *LoginUserRequest) Validate() error {
	if r.Username == "" {
		return errors.New("username is required")
	}

	if r.Password == "" {
		return errors.New("password is required")
	}

	return nil
}


