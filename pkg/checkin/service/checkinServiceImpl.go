package service

import (
	"github.com/jaayroots/habit-tracker-api/entities"
	_checkinRepository "github.com/jaayroots/habit-tracker-api/pkg/checkin/repository"
	_habitRepository "github.com/jaayroots/habit-tracker-api/pkg/habit/repository"
	_userRepository "github.com/jaayroots/habit-tracker-api/pkg/user/repository"
	"github.com/jaayroots/habit-tracker-api/utils"
	"github.com/labstack/echo/v4"

	_checkinMapper "github.com/jaayroots/habit-tracker-api/pkg/checkin/mapper"
	_checkinModel "github.com/jaayroots/habit-tracker-api/pkg/checkin/model"
	_habitException "github.com/jaayroots/habit-tracker-api/pkg/habit/exception"
	_userException "github.com/jaayroots/habit-tracker-api/pkg/user/exception"
	_userModel "github.com/jaayroots/habit-tracker-api/pkg/user/model"
)

type checkinServiceImpl struct {
	habitRepository   _habitRepository.HabitRepository
	checkinRepository _checkinRepository.CheckinRepository
	userRepository    _userRepository.UserRepository
}

func NewCheckinServiceImpl(
	habitRepository _habitRepository.HabitRepository,
	checkinRepository _checkinRepository.CheckinRepository,
	userRepository _userRepository.UserRepository,

) CheckinService {
	return &checkinServiceImpl{
		habitRepository:   habitRepository,
		checkinRepository: checkinRepository,
		userRepository:    userRepository,
	}
}

func (s *checkinServiceImpl) Create(pctx echo.Context, checkinReq *_checkinModel.CheckinReq) (*_checkinModel.CheckinRes, error) {

	val := pctx.Get("user")
	user, ok := val.(*_userModel.UserRes)
	if !ok {
		return nil, _userException.NotFoundUser()
	}

	checkinEntity, err := _checkinMapper.ToCheckinEntity(pctx, checkinReq, 0)
	if err != nil {
		return nil, err
	}

	habit, err := s.habitRepository.FindByIDAndUserID(pctx, checkinEntity.HabitID, uint(user.ID))
	if err != nil {
		return nil, err
	}

	if habit == nil {
		return nil, _habitException.NotFoundHabit()
	}

	checkinEntity, err = s.checkinRepository.Create(pctx, checkinEntity)
	if err != nil {
		return nil, err
	}

	userArr, err := s.prepareUserBlamable(checkinEntity)
	if err != nil {
		return nil, err
	}

	checkinRes := _checkinMapper.ToCheckinRes(checkinEntity, userArr)
	return checkinRes, nil
}

func (s *checkinServiceImpl) prepareUserBlamable(habitEntity *entities.Checkin) ([]*entities.User, error) {
	userIDArray := utils.ExtractAuditUserID(habitEntity)
	userArr, err := s.userRepository.FindByIDs(userIDArray)
	if err != nil {
		return nil, err
	}

	return userArr, nil
}

func (s *checkinServiceImpl) FindAll(pctx echo.Context, checkinSearchReq *_checkinModel.CheckinSearchReq) (*_checkinModel.CheckinSearchRes, error) {

	checkin, total, err := s.checkinRepository.FindAll(pctx, checkinSearchReq)
	if err != nil {
		return nil, err
	}

	userIDArray := utils.ExtractAuditUserIDs(checkin)
	userArr, err := s.userRepository.FindByIDs(userIDArray)
	if err != nil {
		return nil, err
	}

	return _checkinMapper.ToCheckinSearchRes(checkinSearchReq, userArr, checkin, total), nil

}

func (s *checkinServiceImpl) Delete(pctx echo.Context, checkinID uint) (*_checkinModel.CheckinRes, error) {

	checkinEntity, err := s.checkinRepository.Delete(pctx, checkinID)
	if err != nil {
		return nil, err
	}

	userArr, err := s.prepareUserBlamable(checkinEntity)
	if err != nil {
		return nil, err
	}
	checkinRes := _checkinMapper.ToCheckinRes(checkinEntity, userArr)

	return checkinRes, nil
}
