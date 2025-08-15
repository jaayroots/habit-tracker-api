package repository

import (
	"errors"

	databases "github.com/jaayroots/habit-tracker-api/database"
	"github.com/jaayroots/habit-tracker-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_checkinException "github.com/jaayroots/habit-tracker-api/exception/checkin"
	_checkinModel "github.com/jaayroots/habit-tracker-api/model/checkin"
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

type CheckinRepository interface {
	Create(pctx echo.Context, checkin *entities.Checkin) (*entities.Checkin, error)
	FindAll(pctx echo.Context, checkinSearchReq *_checkinModel.CheckinSearchReq) ([]*entities.Checkin, int, error)
	FindByID(pctx echo.Context, checkinID uint) (*entities.Checkin, error)
	Delete(pctx echo.Context, checkinID uint) (*entities.Checkin, error)
	GroupByHabitIDcheckin(pctx echo.Context, habitIDs []uint) ([]*_checkinModel.GroupByHabitIDcheckin, error)
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
		Model(&entities.Checkin{})

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

func (r *checkinRepositoryImpl) Delete(pctx echo.Context, checkinID uint) (*entities.Checkin, error) {

	checkinEntity, err := r.FindByID(pctx, checkinID)
	if err != nil {
		return nil, err
	}

	ctx := pctx.Request().Context()
	err = r.db.Connect().
		WithContext(ctx).
		Delete(checkinEntity).Error

	if err != nil {
		return nil, _checkinException.CannotDeleteCheckin()
	}

	return checkinEntity, nil
}

func (r *checkinRepositoryImpl) FindByID(pctx echo.Context, checkinID uint) (*entities.Checkin, error) {

	checkin := new(entities.Checkin)
	ctx := pctx.Request().Context()
	err := r.db.Connect().
		WithContext(ctx).
		Model(&entities.Checkin{}).
		First(checkin, checkinID).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, _checkinException.NotFoundCheckin()
		}
		return nil, err
	}

	return checkin, nil
}

func (r *checkinRepositoryImpl) searchFilter(query *gorm.DB, checkinSearchReq *_checkinModel.CheckinSearchReq) *gorm.DB {

	query = r.filterHabit(query, checkinSearchReq.Filter)
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

func (r *checkinRepositoryImpl) GroupByHabitIDcheckin(pctx echo.Context, habitIDs []uint) ([]*_checkinModel.GroupByHabitIDcheckin, error) {

	ctx := pctx.Request().Context()

	var checkin []*_checkinModel.GroupByHabitIDcheckin
	query := r.db.Connect().
		WithContext(ctx).
		Model(&entities.Checkin{}).
		Select("habit_id, COUNT(id) as count").
		Group("habit_id").
		Where("habit_id IN ?", habitIDs)

	if err := query.Find(&checkin).Error; err != nil {
		return nil, err
	}

	return checkin, nil
}
