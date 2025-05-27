package mapper

import (
	"github.com/jaayroots/go_base/entities"
	_authUtils "github.com/jaayroots/go_base/pkg/auth/utils"
	_userModel "github.com/jaayroots/go_base/pkg/user/model"
)

func ToUserEntity(userReq *_userModel.UserReq) (*entities.User, error) {
	hashedPassword, err := _authUtils.HashPassword(userReq.Password)
	if err != nil {
		return nil, err
	}

	userEntity := &entities.User{
		Email:     userReq.Email,
		Password:  string(hashedPassword),
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		Avatar:    userReq.Avatar,
	}

	return userEntity, nil
}

func ToUserUpdateEntity(userReq *_userModel.UserUpdateReq) (*entities.User, error) {

	userEntity := &entities.User{
		Email:     userReq.Email,
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
		Avatar:    userReq.Avatar,
	}

	return userEntity, nil
}

func ToUserRes(user *entities.User) *_userModel.UserRes {

	return &_userModel.UserRes{
		ID:        int(user.ID),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Avatar:    user.Avatar,
	}
}
