package mapper

import (
	"github.com/jaayroots/habit-tracker-api/entities"
	"github.com/jaayroots/habit-tracker-api/utils"

	_userModel "github.com/jaayroots/habit-tracker-api/pkg/user/model"
)

func ToUserEntity(userReq *_userModel.UserReq) (*entities.User, error) {
	hashedPassword, err := utils.HashPassword(userReq.Password)
	if err != nil {
		return nil, err
	}

	userEntity := &entities.User{
		Email:     userReq.Email,
		Password:  string(hashedPassword),
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
	}

	return userEntity, nil
}

func ToUserUpdateEntity(userReq *_userModel.UserUpdateReq) (*entities.User, error) {

	userEntity := &entities.User{
		Email:     userReq.Email,
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
	}

	return userEntity, nil
}

func ToUserRes(user *entities.User) *_userModel.UserRes {

	return &_userModel.UserRes{
		ID:        int(user.ID),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}
