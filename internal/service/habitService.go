package service

import (
	"github.com/jaayroots/habit-tracker-api/entities"
	_habitMapper "github.com/jaayroots/habit-tracker-api/mapper/habit"
	_habitModel "github.com/jaayroots/habit-tracker-api/model/habit"
	_repository "github.com/jaayroots/habit-tracker-api/repository"
	"github.com/jaayroots/habit-tracker-api/utils"
	"github.com/labstack/echo/v4"

	_userException "github.com/jaayroots/habit-tracker-api/exception/user"
	_userModel "github.com/jaayroots/habit-tracker-api/model/user"
)

type habitServiceImpl struct {
	habitRepository   _repository.HabitRepository
	userRepository    _repository.UserRepository
	checkinRepository _repository.CheckinRepository
}

func NewHabitServiceImpl(
	habitRepository _repository.HabitRepository,
	userRepository _repository.UserRepository,
	checkinRepository _repository.CheckinRepository,

) HabitService {
	return &habitServiceImpl{
		habitRepository:   habitRepository,
		userRepository:    userRepository,
		checkinRepository: checkinRepository,
	}
}

type HabitService interface {
	Create(pctx echo.Context, habitReq *_habitModel.HabitReq) (*_habitModel.HabitRes, error)
	FindByID(pctx echo.Context, habitID uint) (*_habitModel.HabitRes, error)
	Update(pctx echo.Context, habitID uint, habitReq *_habitModel.HabitReq) (*_habitModel.HabitRes, error)
	Delete(pctx echo.Context, habitID uint) (*_habitModel.HabitRes, error)
	FindAll(pctx echo.Context, habitSearchReq *_habitModel.HabitSearchReq) (*_habitModel.HabitSearchRes, error)
}

func (s *habitServiceImpl) Create(pctx echo.Context, habitReq *_habitModel.HabitReq) (*_habitModel.HabitRes, error) {
	habitEntity, err := _habitMapper.ToHabitEntity(pctx, habitReq, 0)
	if err != nil {
		return nil, err
	}

	val := pctx.Get("user")
	user, ok := val.(*_userModel.UserRes)
	if !ok {
		return nil, _userException.NotFoundUser()
	}

	habitEntity.UserID = user.ID

	habitEntity, err = s.habitRepository.Create(pctx, habitEntity)
	if err != nil {
		return nil, err
	}

	userArr, err := s.prepareUserBlamable(habitEntity)
	if err != nil {
		return nil, err
	}

	habitRes := _habitMapper.ToHabitRes(habitEntity, nil, userArr)
	return habitRes, nil
}

func (s *habitServiceImpl) FindByID(pctx echo.Context, habitID uint) (*_habitModel.HabitRes, error) {

	habitEntity, err := s.habitRepository.FindByID(pctx, habitID)
	if err != nil {
		return nil, err
	}

	userArr, err := s.prepareUserBlamable(habitEntity)
	if err != nil {
		return nil, err
	}

	var habitIDs []uint
	habitIDs = append(habitIDs, habitEntity.ID)

	groupByHabitIDcheckin, err := s.checkinRepository.GroupByHabitIDcheckin(pctx, habitIDs)
	if err != nil {
		return nil, err
	}

	habitRes := _habitMapper.ToHabitRes(habitEntity, groupByHabitIDcheckin, userArr)
	return habitRes, nil
}

func (s *habitServiceImpl) Update(pctx echo.Context, habitID uint, habitReq *_habitModel.HabitReq) (*_habitModel.HabitRes, error) {

	habitEntity, err := _habitMapper.ToHabitEntity(pctx, habitReq, habitID)
	if err != nil {
		return nil, err
	}

	habitEntity, err = s.habitRepository.Update(pctx, habitEntity)
	if err != nil {
		return nil, err
	}

	userArr, err := s.prepareUserBlamable(habitEntity)
	if err != nil {
		return nil, err
	}

	habitRes := _habitMapper.ToHabitRes(habitEntity, nil, userArr)
	return habitRes, nil
}

func (s *habitServiceImpl) Delete(pctx echo.Context, habitID uint) (*_habitModel.HabitRes, error) {

	habitEntity, err := s.habitRepository.Delete(pctx, habitID)
	if err != nil {
		return nil, err
	}

	userArr, err := s.prepareUserBlamable(habitEntity)
	if err != nil {
		return nil, err
	}
	habitRes := _habitMapper.ToHabitRes(habitEntity, nil, userArr)
	return habitRes, nil
}

func (s *habitServiceImpl) prepareUserBlamable(habitEntity *entities.Habit) ([]*entities.User, error) {
	userIDArray := utils.ExtractAuditUserID(habitEntity)
	userArr, err := s.userRepository.FindByIDs(userIDArray)
	if err != nil {
		return nil, err
	}

	return userArr, nil
}

func (s *habitServiceImpl) FindAll(pctx echo.Context, habitSearchReq *_habitModel.HabitSearchReq) (*_habitModel.HabitSearchRes, error) {

	habits, total, err := s.habitRepository.FindAll(pctx, habitSearchReq)
	if err != nil {
		return nil, err
	}

	var habitIDs []uint
	for _, h := range habits {
		habitIDs = append(habitIDs, h.ID)
	}

	groupByHabitIDcheckin, err := s.checkinRepository.GroupByHabitIDcheckin(pctx, habitIDs)
	if err != nil {
		return nil, err
	}

	userIDArray := utils.ExtractAuditUserIDs(habits)
	userArr, err := s.userRepository.FindByIDs(userIDArray)
	if err != nil {
		return nil, err
	}

	return _habitMapper.ToHabitSearchRes(habitSearchReq, userArr, habits, groupByHabitIDcheckin, total), nil

}
