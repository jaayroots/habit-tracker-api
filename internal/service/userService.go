package service

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	databases "github.com/habit-tracker-api/database"
	_exceptionType "github.com/habit-tracker-api/enums/exception"
	_exception "github.com/habit-tracker-api/exception"
	_userMapper "github.com/habit-tracker-api/mapper/user"
	_userContact "github.com/habit-tracker-api/mapper/userContact"
	_userContactMapper "github.com/habit-tracker-api/mapper/userContact"
	_userModel "github.com/habit-tracker-api/model/user"
	_repository "github.com/habit-tracker-api/repository"
)

type userServiceImpl struct {
	db                    databases.Database
	userRepository        _repository.UserRepository
	userContactRepository _repository.UserContactRepository
}

func NewUserService(
	db databases.Database,
	userRepository _repository.UserRepository,
	userContactRepository _repository.UserContactRepository,
) UserService {
	return &userServiceImpl{
		db:                    db,
		userRepository:        userRepository,
		userContactRepository: userContactRepository,
	}
}

type UserService interface {
	FindByID(userID uuid.UUID) (*_userModel.UserRes, error)
	Update(userID uuid.UUID, userUpdateReq *_userModel.UserUpdateReq) error
	Delete(userID uuid.UUID) error
}

func (s *userServiceImpl) FindByID(userID uuid.UUID) (*_userModel.UserRes, error) {
	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, _exception.Handle("not found user", _exceptionType.Info)
	}

	userContact, err := s.userContactRepository.FindByUserID(userID)
	if err != nil {
		return nil, err
	}
	userContactRes := _userContact.ToUserContactRes(userContact)

	userRes := _userMapper.ToUserRes(user, userContactRes)
	return userRes, nil
}

func (s *userServiceImpl) Update(userID uuid.UUID, userUpdateReq *_userModel.UserUpdateReq) error {

	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		return err
	}

	if user == nil {
		return _exception.Handle("not found user", _exceptionType.Info)
	}

	userContactCreateEntity, _, err := _userContactMapper.ToUserContactCreateEntity(user.ID, userUpdateReq.UserContact)
	if err != nil {
		return err
	}

	userContactUpdateEntity, _, err := _userContactMapper.ToUserContactUpdateEntity(user.ID, userUpdateReq.UserContact)
	if err != nil {
		return err
	}

	userContactDeleteEntity, _, err := _userContactMapper.ToUserContactDeleteEntity(user.ID, userUpdateReq.UserContact)
	if err != nil {
		return err
	}

	err = s.db.Connect().Transaction(func(tx *gorm.DB) error {
		if err := s.userRepository.Update(tx, user); err != nil {
			return err
		}

		if err := s.userContactRepository.Create(tx, userContactCreateEntity); err != nil {
			return err
		}

		if err := s.userContactRepository.Update(tx, userContactUpdateEntity); err != nil {
			return err
		}

		if err := s.userContactRepository.Delete(tx, userContactDeleteEntity); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (s *userServiceImpl) Delete(userID uuid.UUID) error {

	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return _exception.Handle("not found user", _exceptionType.Info)
	}

	err = s.db.Connect().Transaction(func(tx *gorm.DB) error {
		if err := s.userRepository.Delete(tx, user.ID); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
