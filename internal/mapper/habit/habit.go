package mapper

import (
	"github.com/jaayroots/habit-tracker-api/entities"
	"github.com/jaayroots/habit-tracker-api/enums"
	_exceptionType "github.com/jaayroots/habit-tracker-api/enums/exception"
	_exception "github.com/jaayroots/habit-tracker-api/exception"
	_checkinModel "github.com/jaayroots/habit-tracker-api/model/checkin"
	_habitModel "github.com/jaayroots/habit-tracker-api/model/habit"
	"github.com/jaayroots/habit-tracker-api/utils"

	"github.com/labstack/echo/v4"
)

func ToHabitEntity(pctx echo.Context, habitReq *_habitModel.HabitReq, habitID uint) (*entities.Habit, error) {

	frequency := enums.Frequency(habitReq.Frequency)
	if !enums.IsValidFrequency(int(frequency)) {
		return nil, _exception.Handle("frequency invalid", _exceptionType.Info)
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

func ToHabitRes(habit *entities.Habit, checkinGroup []*_checkinModel.GroupByHabitIDcheckin, user []*entities.User) *_habitModel.HabitRes {

	userMap := utils.MapperByID(user)
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

	habitMap := make(map[uint]uint)
	for _, c := range checkinGroup {
		habitMap[c.HabitID] = uint(c.Count)
	}

	return &_habitModel.HabitRes{
		ID:          int(habit.ID),
		Title:       habit.Title,
		Description: habit.Description,
		Frequency:   int(habit.Frequency),
		TargetCount: int(habit.TargetCount),
		Checkin:     int(habitMap[habit.ID]),
		CreatedAt:   habit.CreatedAt.Unix(),
		UpdatedAt:   habit.UpdatedAt.Unix(),
		CreatedBy:   createdBy,
		UpdatedBy:   updatedBy,
	}
}

func ToHabitSearchRes(
	habitSearchReq *_habitModel.HabitSearchReq,
	user []*entities.User,
	habits []*entities.Habit,
	checkinGroup []*_checkinModel.GroupByHabitIDcheckin,
	total int,
) *_habitModel.HabitSearchRes {

	habitResList := make([]*_habitModel.HabitRes, 0, len(habits))
	for _, habit := range habits {
		habitResList = append(habitResList, ToHabitRes(habit, checkinGroup, user))
	}

	_, _, totalPage := utils.PaginateCalculate(habitSearchReq.Page, habitSearchReq.Limit, total)
	return &_habitModel.HabitSearchRes{
		Item: habitResList,
		Paginate: _habitModel.PaginateResult{
			Page:      int64(habitSearchReq.Page),
			TotalPage: int64(totalPage),
			Total:     int64(total),
		},
	}
}
