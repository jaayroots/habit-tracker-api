package service

import (
	_authModel "github.com/jaayroots/go_base/pkg/auth/model"
	_userModel "github.com/jaayroots/go_base/pkg/user/model"
)

type AuthService interface {
	Login(loginReq *_authModel.LoginReq) (*_authModel.LoginRes, error)
	Logout(userID uint) error
	Register(userReq *_userModel.UserReq) (*_userModel.UserRes, error)
	Authorizing(token string) (*_authModel.LoginRes, bool, error)
	Refreash(ipAddress string, userID uint) (*_authModel.LoginRes, error)
}
