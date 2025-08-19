package mapper

import (
	"github.com/habit-tracker-api/entities"
	_checkinModel "github.com/habit-tracker-api/model/checkin"
	"github.com/habit-tracker-api/utils"

	"github.com/labstack/echo/v4"
)

func ToCheckinEntity(pctx echo.Context, checkinReq *_checkinModel.CheckinReq, checkinID uint) (*entities.Checkin, error) {

	checkinEntity := &entities.Checkin{
		ID:      checkinID,
		HabitID: uint(checkinReq.HabitID),
	}

	return checkinEntity, nil
}

func ToCheckinRes(checkin *entities.Checkin, user []*entities.User) *_checkinModel.CheckinRes {

	userMap := utils.MapperByID(user)
	createdBy := func() *string {
		if user, ok := userMap[checkin.CreatedBy]; ok {
			fullName := user.FirstName + " " + user.LastName
			return &fullName
		}
		return nil
	}()

	updatedBy := func() *string {
		if user, ok := userMap[checkin.UpdatedBy]; ok {
			fullName := user.FirstName + " " + user.LastName
			return &fullName
		}
		return nil
	}()

	return &_checkinModel.CheckinRes{
		ID:        int(checkin.ID),
		HabitID:   int(checkin.HabitID),
		CreatedAt: checkin.CreatedAt.Unix(),
		UpdatedAt: checkin.UpdatedAt.Unix(),
		CreatedBy: createdBy,
		UpdatedBy: updatedBy,
	}
}

func ToCheckinSearchRes(
	checkinSearchReq *_checkinModel.CheckinSearchReq,
	user []*entities.User,
	checkins []*entities.Checkin,
	total int,
) *_checkinModel.CheckinSearchRes {

	checkinResList := make([]*_checkinModel.CheckinRes, 0, len(checkins))
	for _, checkin := range checkins {
		checkinResList = append(checkinResList, ToCheckinRes(checkin, user))
	}

	_, _, totalPage := utils.PaginateCalculate(checkinSearchReq.Page, checkinSearchReq.Limit, total)
	return &_checkinModel.CheckinSearchRes{
		Item: checkinResList,
		Paginate: _checkinModel.PaginateResult{
			Page:      int64(checkinSearchReq.Page),
			TotalPage: int64(totalPage),
			Total:     int64(total),
		},
	}
}
