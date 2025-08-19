package repository

import (
	"errors"

	"github.com/google/uuid"
	databases "github.com/habit-tracker-api/database"
	"github.com/habit-tracker-api/entities"
	_exceptionType "github.com/habit-tracker-api/enums/exception"
	_exception "github.com/habit-tracker-api/exception"
	_habitModel "github.com/habit-tracker-api/model/habit"
	"github.com/habit-tracker-api/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type habitRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

type CheckinGroup struct {
	HabitID uint
	Count   int64
}

type HabitRepository interface {
	Create(pctx echo.Context, habit *entities.Habit) (*entities.Habit, error)
	FindByID(pctx echo.Context, habitID uint) (*entities.Habit, error)
	Update(pctx echo.Context, habit *entities.Habit) (*entities.Habit, error)
	Delete(pctx echo.Context, habitID uint) (*entities.Habit, error)
	FindAll(pctx echo.Context, habitSearchReq *_habitModel.HabitSearchReq) ([]*entities.Habit, int, error)
	FindByIDAndUserID(pctx echo.Context, habitID uint, userID uuid.UUID) (*entities.Habit, error)
}

func NewHabitRepositoryImpl(db databases.Database, logger echo.Logger) HabitRepository {
	return &habitRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *habitRepositoryImpl) Create(pctx echo.Context, habit *entities.Habit) (*entities.Habit, error) {

	habitEntity := new(entities.Habit)
	ctx := pctx.Request().Context()

	err := r.db.Connect().
		WithContext(ctx).
		Create(habit).
		Scan(habitEntity).
		Error

	if err != nil {
		return nil, _exception.Handle("cannot create habbit", _exceptionType.Info)
	}
	return habit, nil
}

func (r *habitRepositoryImpl) FindByID(pctx echo.Context, habitID uint) (*entities.Habit, error) {

	habit := new(entities.Habit)
	ctx := pctx.Request().Context()
	err := r.db.Connect().
		WithContext(ctx).
		Model(&entities.Habit{}).
		First(habit, habitID).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, _exception.Handle("not found habbit", _exceptionType.Info)
		}
		return nil, err
	}

	return habit, nil
}

func (r *habitRepositoryImpl) Update(pctx echo.Context, habit *entities.Habit) (*entities.Habit, error) {
	_, err := r.FindByID(pctx, habit.ID)
	if err != nil {
		return nil, err
	}

	habitEntity := new(entities.Habit)
	ctx := pctx.Request().Context()
	err = r.db.Connect().WithContext(ctx).
		Updates(habit).
		Scan(habitEntity).
		Error
	if err != nil {
		return nil, _exception.Handle("cannot update habbit", _exceptionType.Info)
	}

	return habitEntity, nil
}

func (r *habitRepositoryImpl) Delete(pctx echo.Context, habitID uint) (*entities.Habit, error) {

	habitEntity, err := r.FindByID(pctx, habitID)
	if err != nil {
		return nil, err
	}

	ctx := pctx.Request().Context()
	err = r.db.Connect().
		WithContext(ctx).
		Delete(habitEntity).Error

	if err != nil {
		return nil, _exception.Handle("cannot delete habbit", _exceptionType.Info)
	}

	return habitEntity, nil
}

func (r *habitRepositoryImpl) FindAll(pctx echo.Context, habitSearchReq *_habitModel.HabitSearchReq) ([]*entities.Habit, int, error) {

	var habit []*entities.Habit
	var total int64
	ctx := pctx.Request().Context()
	query := r.db.Connect().
		WithContext(ctx).
		Model(&entities.Habit{})

	offset, limit, _ := utils.PaginateCalculate(habitSearchReq.Page, habitSearchReq.Limit, 0)
	query = r.searchFilter(query, habitSearchReq)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&habit).Error; err != nil {
		return nil, 0, err
	}

	if len(habit) == 0 {
		return nil, 0, nil
	}

	return habit, int(total), nil
}

func (r *habitRepositoryImpl) searchFilter(query *gorm.DB, habitSearchReq *_habitModel.HabitSearchReq) *gorm.DB {

	query = r.filterTitle(query, habitSearchReq.Filter)
	query = r.filterDescription(query, habitSearchReq.Filter)
	query = r.filterFrequency(query, habitSearchReq.Filter)
	query = r.filterTargetCount(query, habitSearchReq.Filter)
	return query
}

func (r *habitRepositoryImpl) filterTitle(query *gorm.DB, habitFilterReq _habitModel.HabitFilterReq) *gorm.DB {

	title := habitFilterReq.Title
	if title == nil || *title == "" {
		return query
	}

	query = query.Where("title ILIKE ?", "%"+*title+"%")
	return query
}

func (r *habitRepositoryImpl) filterDescription(query *gorm.DB, habitFilterReq _habitModel.HabitFilterReq) *gorm.DB {

	description := habitFilterReq.Description
	if description == nil || *description == "" {
		return query
	}

	query = query.Where("description ILIKE ?", "%"+*description+"%")
	return query
}

func (r *habitRepositoryImpl) filterFrequency(query *gorm.DB, habitFilterReq _habitModel.HabitFilterReq) *gorm.DB {

	frequency := habitFilterReq.Frequency
	if frequency == nil {
		return query
	}

	query = query.Where("frequency = ?", *frequency)
	return query
}

func (r *habitRepositoryImpl) filterTargetCount(query *gorm.DB, habitFilterReq _habitModel.HabitFilterReq) *gorm.DB {

	targetCount := habitFilterReq.TargetCount
	if targetCount == nil {
		return query
	}

	query = query.Where("target_count = ?", *targetCount)
	return query
}

func (r *habitRepositoryImpl) FindByIDAndUserID(pctx echo.Context, habitID uint, userID uuid.UUID) (*entities.Habit, error) {
	ctx := pctx.Request().Context()
	var habit entities.Habit

	err := r.db.Connect().
		WithContext(ctx).
		Where("id = ? AND user_id = ?", habitID, userID).
		First(&habit).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, _exception.Handle("not found habbit", _exceptionType.Info)
		}
		return nil, err
	}

	return &habit, nil
}
