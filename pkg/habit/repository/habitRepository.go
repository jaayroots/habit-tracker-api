package repository

import (
	"github.com/jaayroots/habit-tracker-api/entities"
	"github.com/labstack/echo/v4"
)

type HabitRepository interface {
	Create(pctx echo.Context, habit *entities.Habit) (*entities.Habit, error)
	FindByID(pctx echo.Context, habitID uint) (*entities.Habit, error)
	Update(pctx echo.Context, habit *entities.Habit) (*entities.Habit, error)
	Delete(pctx echo.Context, habitID uint) (*entities.Habit, error)
}
