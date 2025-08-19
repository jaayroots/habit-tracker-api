package mapper

import (
	"github.com/google/uuid"
	"github.com/habit-tracker-api/entities"
	"github.com/habit-tracker-api/utils"

	_userModel "github.com/habit-tracker-api/model/user"
	_userContactModel "github.com/habit-tracker-api/model/userContact"
)

func ToUserEntity(userReq *_userModel.UserReq) (*entities.User, error) {
	hashedPassword, err := utils.HashPassword(userReq.Password)
	if err != nil {
		return nil, err
	}

	userEntity := &entities.User{
		ID:        uuid.New(),
		Email:     userReq.Email,
		Password:  string(hashedPassword),
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
	}

	return userEntity, nil
}

func ToUserUpdateEntity(userID uuid.UUID, userReq *_userModel.UserUpdateReq) (*entities.User, error) {

	userEntity := &entities.User{
		ID:        userID,
		Email:     userReq.Email,
		FirstName: userReq.FirstName,
		LastName:  userReq.LastName,
	}

	return userEntity, nil
}

func ToUserLoginRes(user *entities.User) *_userModel.UserRes {

	return &_userModel.UserRes{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func ToUserRes(user *entities.User, userContactRes []*_userContactModel.UserContactRes) *_userModel.UserRes {

	return &_userModel.UserRes{
		ID:          user.ID,
		Email:       user.Email,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		UserContact: userContactRes,
	}
}
