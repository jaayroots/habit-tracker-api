package repository

import (
	"github.com/jaayroots/habit-tracker-api/entities"
	_habitModel "github.com/jaayroots/habit-tracker-api/pkg/habit/model"
	"github.com/labstack/echo/v4"
)

type HabitRepository interface {
	Create(pctx echo.Context, habit *entities.Habit) (*entities.Habit, error)
	FindByID(pctx echo.Context, habitID uint) (*entities.Habit, error)
	Update(pctx echo.Context, habit *entities.Habit) (*entities.Habit, error)
	Delete(pctx echo.Context, habitID uint) (*entities.Habit, error)
	FindAll(pctx echo.Context, habitSearchReq *_habitModel.HabitSearchReq) ([]*entities.Habit, int, error)
}
