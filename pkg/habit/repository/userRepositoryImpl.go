package repository

import (
	"errors"

	databases "github.com/jaayroots/habit-tracker-api/database"
	"github.com/jaayroots/habit-tracker-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_habitException "github.com/jaayroots/habit-tracker-api/pkg/habit/exception"
	_habitModel "github.com/jaayroots/habit-tracker-api/pkg/habit/model"
	"github.com/jaayroots/habit-tracker-api/utils"
)

type habitRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
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
		return nil, _habitException.CannotCreateHabit()
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
			return nil, _habitException.NotFoundHabit()
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
		return nil, _habitException.CannotUpdateHabit()
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
		return nil, _habitException.CannotDeleteHabit()
	}

	return habitEntity, nil
}

func (r *habitRepositoryImpl) FindAll(pctx echo.Context, habitSearchReq *_habitModel.HabitSearchReq) ([]*entities.Habit, int, error) {

	var habit []*entities.Habit
	var total int64
	ctx := pctx.Request().Context()
	query := r.db.Connect().WithContext(ctx).Model(&entities.Habit{})

	offset, limit, _ := utils.PaginateCalculate(habitSearchReq.Page, habitSearchReq.Limit, 0)
	query = r.searchFilter(query, habitSearchReq)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&habit).Error; err != nil {
		return nil, 0, err
	}

	return habit, int(total), nil
}

func (r *habitRepositoryImpl) searchFilter(query *gorm.DB, habitSearchReq *_habitModel.HabitSearchReq) *gorm.DB {

	query = r.filterTitle(query, habitSearchReq.Filter)
	query = r.filterDescription(query, habitSearchReq.Filter)
	query = r.filterFrequency(query, habitSearchReq.Filter)
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
