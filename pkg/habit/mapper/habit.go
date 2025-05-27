package mapper

import (
	"github.com/jaayroots/habit-tracker-api/entities"
	"github.com/jaayroots/habit-tracker-api/enums"
	_habitException "github.com/jaayroots/habit-tracker-api/pkg/habit/exception"
	_habitModel "github.com/jaayroots/habit-tracker-api/pkg/habit/model"

	"github.com/labstack/echo/v4"

	_utils "github.com/jaayroots/habit-tracker-api/utils"
)

func ToHabitEntity(pctx echo.Context, habitReq *_habitModel.HabitReq, habitID uint) (*entities.Habit, error) {

	frequency := enums.Frequency(habitReq.Frequency)
	if !enums.IsValidFrequency(int(frequency)) {
		return nil, _habitException.FrequencyInvalid()
	}

	habitEntity := &entities.Habit{
		ID:          habitID,
		Title:       habitReq.Title,
		Description: habitReq.Description,
		Frequency:   frequency,
		TargetCount: uint(habitReq.TargetCount),
	}

	return habitEntity, nil
}

func ToHabitRes(habit *entities.Habit, user []*entities.User) *_habitModel.HabitRes {

	userMap := _utils.MapperByID(user)
	createdBy := func() *string {
		if user, ok := userMap[habit.CreatedBy]; ok {
			fullName := user.FirstName + " " + user.LastName
			return &fullName
		}
		return nil
	}()

	updatedBy := func() *string {
		if user, ok := userMap[habit.UpdatedBy]; ok {
			fullName := user.FirstName + " " + user.LastName
			return &fullName
		}
		return nil
	}()

	return &_habitModel.HabitRes{
		ID:          int(habit.ID),
		Title:       habit.Title,
		Description: habit.Description,
		Frequency:   int(habit.Frequency),
		TargetCount: int(habit.TargetCount),
		CreatedAt:   habit.CreatedAt.Unix(),
		UpdatedAt:   habit.UpdatedAt.Unix(),
		CreatedBy:   createdBy,
		UpdatedBy:   updatedBy,
	}
}
