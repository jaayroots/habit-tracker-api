package service

import (
	"time"

	"github.com/google/uuid"
	"github.com/habit-tracker-api/config"
	databases "github.com/habit-tracker-api/database"
	"github.com/habit-tracker-api/utils"
	"gorm.io/gorm"

	_exceptionType "github.com/habit-tracker-api/enums/exception"
	_exception "github.com/habit-tracker-api/exception"
	_authMapper "github.com/habit-tracker-api/mapper/auth"
	_userMapper "github.com/habit-tracker-api/mapper/user"
	_userContactMapper "github.com/habit-tracker-api/mapper/userContact"
	_authModel "github.com/habit-tracker-api/model/auth"
	_userModel "github.com/habit-tracker-api/model/user"
	_repository "github.com/habit-tracker-api/repository"
)

type authServiceImpl struct {
	db                    databases.Database
	userRepository        _repository.UserRepository
	userContactRepository _repository.UserContactRepository
	sessionRepository     _repository.SessionRepository
}

func NewAuthServiceImpl(
	db databases.Database,
	userRepository _repository.UserRepository,
	userContactRepository _repository.UserContactRepository,
	sessionRepository _repository.SessionRepository,
) AuthService {
	return &authServiceImpl{
		db:                    db,
		userRepository:        userRepository,
		userContactRepository: userContactRepository,
		sessionRepository:     sessionRepository,
	}
}

type AuthService interface {
	Login(loginReq *_authModel.LoginReq) (*_authModel.LoginRes, error)
	Logout(userID uuid.UUID) error
	Register(userReq *_userModel.UserReq) (*_userModel.UserRes, error)
	Authorizing(token string) (*_authModel.LoginRes, bool, error)
	Refreash(ipAddress string, userID uuid.UUID) (*_authModel.LoginRes, error)
}

func (s *authServiceImpl) Register(userReq *_userModel.UserReq) (*_userModel.UserRes, error) {

	userEntity, err := _userMapper.ToUserEntity(userReq)
	if err != nil {
		return nil, err
	}

	userContactEntity, _, err := _userContactMapper.ToUserContactCreateEntity(userEntity.ID, userReq.UserContact)
	if err != nil {
		return nil, err
	}

	user, err := s.userRepository.FindByEmail(userEntity.Email)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, _exception.Handle("email is already", _exceptionType.Info)
	}

	err = s.db.Connect().Transaction(func(tx *gorm.DB) error {
		if err := s.userRepository.Create(tx, userEntity); err != nil {
			return err
		}

		if len(userContactEntity) != 0 {
			if err := s.userContactRepository.Create(tx, userContactEntity); err != nil {
				return err
			}
		}

		user = userEntity
		return nil
	})

	if err != nil {
		return nil, err
	}

	userRes := _userMapper.ToUserLoginRes(user)
	return userRes, nil
}

func (s *authServiceImpl) Login(loginReq *_authModel.LoginReq) (*_authModel.LoginRes, error) {

	user, err := s.userRepository.FindByEmail(loginReq.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, _exception.Handle("not found user", _exceptionType.Info)
	}

	isValid := utils.CheckPasswordHash(loginReq.Password, user.Password)
	if !isValid {
		return nil, _exception.Handle("authentication failed", _exceptionType.Info)
	}

	expToken := config.ConfigGetting().Security.ExpiredToken
	if expToken == 0 {
		expToken = 24
	}

	token, exp, err := utils.HashToken(_userMapper.ToUserLoginRes(user), expToken)
	if err != nil {
		return nil, err
	}

	session := _authMapper.ToSessionEntity(user, token, exp, loginReq.IpAddress)

	err = s.db.Connect().Transaction(func(tx *gorm.DB) error {

		if err := s.sessionRepository.Delete(tx, session.UserID); err != nil {
			return err
		}

		if err := s.sessionRepository.Create(tx, session); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return _authMapper.ToAuthRes(user, token), nil
}

func (s *authServiceImpl) Logout(userID uuid.UUID) error {

	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		return err
	}

	if user == nil {
		return _exception.Handle("not found user", _exceptionType.Info)
	}

	err = s.db.Connect().Transaction(func(tx *gorm.DB) error {

		if err := s.sessionRepository.Delete(tx, userID); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *authServiceImpl) Authorizing(token string) (*_authModel.LoginRes, bool, error) {

	isTokenExpSoon := false
	session, err := s.sessionRepository.FindByToken(token)
	if err != nil {
		return nil, isTokenExpSoon, err
	}
	if session == nil {
		return nil, isTokenExpSoon, _exception.Handle("token invalid", _exceptionType.Info)
	}

	if time.Now().After(session.ExpiresAt) {
		return nil, isTokenExpSoon, _exception.Handle("token invalid", _exceptionType.Info)
	}

	user, err := s.userRepository.FindByID(session.UserID)
	if err != nil {
		return nil, isTokenExpSoon, err
	}
	if user == nil {
		return nil, isTokenExpSoon, _exception.Handle("token invalid", _exceptionType.Info)
	}

	refreshTokenBefore := config.ConfigGetting().Security.RefreshToken
	if time.Now().After(session.ExpiresAt.Add(-time.Duration(refreshTokenBefore) * time.Minute)) {
		isTokenExpSoon = true
	}

	return _authMapper.ToAuthRes(user, token), isTokenExpSoon, nil
}

func (s *authServiceImpl) Refreash(ipAddress string, userID uuid.UUID) (*_authModel.LoginRes, error) {

	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, _exception.Handle("not found user", _exceptionType.Info)
	}

	err = s.Logout(user.ID)
	if err != nil {
		return nil, err
	}

	expToken := config.ConfigGetting().Security.ExpiredToken
	if expToken == 0 {
		expToken = 24
	}

	token, exp, err := utils.HashToken(_userMapper.ToUserLoginRes(user), expToken)
	if err != nil {
		return nil, err
	}

	session := _authMapper.ToSessionEntity(user, token, exp, ipAddress)
	err = s.db.Connect().Transaction(func(tx *gorm.DB) error {
		if err := s.sessionRepository.Create(tx, session); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return _authMapper.ToAuthRes(user, token), nil

}
