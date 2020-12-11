package models

import (
	"github.com/google/uuid"
)

type LoginUserData struct {
	User struct {
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"user"`
}

type RegisterUserData struct {
	User struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	} `json:"user"`
}

type Id struct {
	id uuid.UUID
}
