package mapper

import (
	"time"

	"github.com/jaayroots/habit-tracker-api/entities"
	_authModel "github.com/jaayroots/habit-tracker-api/model/auth"
	_userModel "github.com/jaayroots/habit-tracker-api/model/user"
)

func ToAuthRes(user *entities.User, token string) *_authModel.LoginRes {

	return &_authModel.LoginRes{
		Token: token,
		User: &_userModel.UserRes{
			ID:        user.ID,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		},
	}
}

func ToSessionEntity(user *entities.User, token string, exp time.Time, ipAddress string) *entities.Session {

	userEntity := &entities.Session{
		UserID:    user.ID,
		Token:     token,
		ExpiresAt: exp,
		IpAddress: ipAddress,
	}

	return userEntity
}
