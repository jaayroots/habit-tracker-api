package service

import (
	_userModel "github.com/jaayroots/go_base/pkg/user/model"
)

type UserService interface {
	FindByID(userID uint) (*_userModel.UserRes, error)
	Update(userID uint, userUpdateReq *_userModel.UserUpdateReq) error
	Delete(userID uint) error
}
