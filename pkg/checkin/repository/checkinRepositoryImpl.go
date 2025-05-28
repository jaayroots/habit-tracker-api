package repository

import (
	databases "github.com/jaayroots/habit-tracker-api/database"
	"github.com/jaayroots/habit-tracker-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_checkinException "github.com/jaayroots/habit-tracker-api/pkg/checkin/exception"
	_checkinModel "github.com/jaayroots/habit-tracker-api/pkg/checkin/model"
	"github.com/jaayroots/habit-tracker-api/utils"
)

type checkinRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewCheckinRepositoryImpl(db databases.Database, logger echo.Logger) CheckinRepository {
	return &checkinRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *checkinRepositoryImpl) Create(pctx echo.Context, checkin *entities.Checkin) (*entities.Checkin, error) {

	checkinEntity := new(entities.Checkin)
	ctx := pctx.Request().Context()

	err := r.db.Connect().
		WithContext(ctx).
		Create(checkin).
		Scan(checkinEntity).
		Error

	if err != nil {
		return nil, _checkinException.CannotCreateCheckin()
	}
	return checkin, nil
}

func (r *checkinRepositoryImpl) FindAll(pctx echo.Context, checkinSearchReq *_checkinModel.CheckinSearchReq) ([]*entities.Checkin, int, error) {
	var checkin []*entities.Checkin
	var total int64
	ctx := pctx.Request().Context()

	query := r.db.Connect().
		WithContext(ctx).
		Model(&entities.Checkin{}).
		Joins("JOIN habits h ON checkins.habit_id = h.id").
		Preload("Habit")

	offset, limit, _ := utils.PaginateCalculate(checkinSearchReq.Page, checkinSearchReq.Limit, 0)
	query = r.searchFilter(query, checkinSearchReq)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Offset(offset).Limit(limit).Order("created_at DESC").Find(&checkin).Error; err != nil {
		return nil, 0, err
	}

	return checkin, int(total), nil
}

func (r *checkinRepositoryImpl) searchFilter(query *gorm.DB, checkinSearchReq *_checkinModel.CheckinSearchReq) *gorm.DB {

	query = r.filterHabit(query, checkinSearchReq.Filter)
	query = r.filterFrequency(query, checkinSearchReq.Filter)
	return query
}

func (r *checkinRepositoryImpl) filterHabit(query *gorm.DB, habitFilterReq _checkinModel.CheckinFilterReq) *gorm.DB {

	HabitID := habitFilterReq.HabitID
	if HabitID == nil || *HabitID == 0 {
		return query
	}

	query = query.Where("checkins.habit_id = ?", HabitID)
	return query
}

func (r *checkinRepositoryImpl) filterFrequency(query *gorm.DB, habitFilterReq _checkinModel.CheckinFilterReq) *gorm.DB {

	Frequency := habitFilterReq.Frequency
	if Frequency == nil || *Frequency == 0 {
		return query
	}

	query = query.Where("h.frequency = ?", Frequency)
	return query
}
