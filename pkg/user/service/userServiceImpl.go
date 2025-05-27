package service

import (
	_userException "github.com/jaayroots/go_base/pkg/user/exception"
	_userMapper "github.com/jaayroots/go_base/pkg/user/mapper"
	_userModel "github.com/jaayroots/go_base/pkg/user/model"
	_userRepository "github.com/jaayroots/go_base/pkg/user/repository"
)

type userServiceImpl struct {
	userRepository _userRepository.UserRepository
}

func NewUserServiceImpl(
	userRepository _userRepository.UserRepository,
) UserService {
	return &userServiceImpl{userRepository}
}

func (s *userServiceImpl) FindByID(userID uint) (*_userModel.UserRes, error) {

	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, _userException.NotFoundUser()
	}

	userRes := _userMapper.ToUserRes(user)
	return userRes, nil
}

func (s *userServiceImpl) Update(userID uint, userUpdateReq *_userModel.UserUpdateReq) error {
	userEntity, err := _userMapper.ToUserUpdateEntity(userUpdateReq)
	if err != nil {
		return err
	}

	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return _userException.NotFoundUser()
	}

	_, err = s.userRepository.Update(userID, userEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *userServiceImpl) Delete(userID uint) error {

	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return _userException.NotFoundUser()
	}

	err = s.userRepository.Delete(userID)
	if err != nil {
		return err
	}

	return nil
}
