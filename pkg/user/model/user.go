package model

import (
	"time"
)

type (
	User struct {
		ID        int       `json:"id"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		Avatar    string    `json:"avatar"`
		CreatedAt time.Time `json:"createdAt"`
	}

	UserReq struct {
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required,min=2,max=64"`
		FirstName string `json:"first_name" validate:"required,min=2,max=100"`
		LastName  string `json:"last_name" validate:"required,min=2,max=100"`
		Avatar    string `json:"avatar"`
	}

	UserUpdateReq struct {
		Email     string `json:"email" validate:"required,email"`
		FirstName string `json:"first_name" validate:"required,min=2,max=100"`
		LastName  string `json:"last_name" validate:"required,min=2,max=100"`
		Avatar    string `json:"avatar"`
	}

	UserRes struct {
		ID        int    `json:"id"`
		Email     string `json:"email"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Avatar    string `json:"avatar"`
	}
)
