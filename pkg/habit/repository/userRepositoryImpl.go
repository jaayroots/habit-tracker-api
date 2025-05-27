package repository

import (
	"errors"

	databases "github.com/jaayroots/habit-tracker-api/database"
	"github.com/jaayroots/habit-tracker-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_habitException "github.com/jaayroots/habit-tracker-api/pkg/habit/exception"
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
