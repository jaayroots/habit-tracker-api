package model

import (
	"time"

	"github.com/google/uuid"
	model "github.com/habit-tracker-api/model/userContact"
)

type (
	User struct {
		ID        uuid.UUID `json:"id"`
		Email     string    `json:"email"`
		Password  string    `json:"password"`
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		CreatedAt time.Time `json:"createdAt"`
	}

	UserReq struct {
		Email       string                  `json:"email" validate:"required,email"`
		Password    string                  `json:"password" validate:"required,min=2,max=64"`
		FirstName   string                  `json:"first_name" validate:"required,min=2,max=100"`
		LastName    string                  `json:"last_name" validate:"required,min=2,max=100"`
		UserContact []*model.UserContactReq `json:"user_contact"`
	}

	UserUpdateReq struct {
		Email       string                  `json:"email" validate:"required,email"`
		FirstName   string                  `json:"first_name" validate:"required,min=2,max=100"`
		LastName    string                  `json:"last_name" validate:"required,min=2,max=100"`
		UserContact []*model.UserContactReq `json:"user_contact"`
	}

	UserRes struct {
		ID          uuid.UUID               `json:"id"`
		Email       string                  `json:"email"`
		FirstName   string                  `json:"first_name"`
		LastName    string                  `json:"last_name"`
		UserContact []*model.UserContactRes `json:"user_contact,omitempty"`
	}
)
