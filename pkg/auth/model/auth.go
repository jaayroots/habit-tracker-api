package model

import (
	_userModel "github.com/jaayroots/habit-tracker-api/pkg/user/model"
)

type (
	LoginReq struct {
		Email     string `json:"email" validate:"required,email"`
		Password  string `json:"password" validate:"required,min=2,max=64"`
		IpAddress string
	}

	LoginRes struct {
		Token string              `json:"token"`
		User  *_userModel.UserRes `json:"user"`
	}
)
