package repository

import (
	"errors"

	"github.com/google/uuid"
	databases "github.com/jaayroots/habit-tracker-api/database"
	"github.com/jaayroots/habit-tracker-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_exceptionType "github.com/jaayroots/habit-tracker-api/enums/exception"
	_exception "github.com/jaayroots/habit-tracker-api/exception"
)

type userRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewUserRepository(db databases.Database, logger echo.Logger) UserRepository {
	return &userRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

type UserRepository interface {
	Create(tx *gorm.DB, user *entities.User) error
	FindByEmail(email string) (*entities.User, error)
	FindByID(userID uuid.UUID) (*entities.User, error)
	FindByIDs(userIDs []uuid.UUID) ([]*entities.User, error)
	Update(tx *gorm.DB, user *entities.User) error
	Delete(tx *gorm.DB, userID uuid.UUID) error
}

func (r *userRepositoryImpl) Create(tx *gorm.DB, user *entities.User) error {
	err := tx.Create(user).Error
	if err != nil {
		return _exception.Handle("cannot create user", _exceptionType.Info)
	}
	return err
}

func (r *userRepositoryImpl) FindByEmail(email string) (*entities.User, error) {

	userEntity := new(entities.User)

	err := r.db.Connect().
		Where("email = ? and is_deleted = ?", email, false).
		First(userEntity).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return userEntity, nil

}

func (r *userRepositoryImpl) FindByID(userID uuid.UUID) (*entities.User, error) {

	user := new(entities.User)

	err := r.db.Connect().
		Where("id = ? and is_deleted = ?", userID, false).
		First(user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepositoryImpl) FindByIDs(userIDs []uuid.UUID) ([]*entities.User, error) {

	if len(userIDs) == 0 {
		return []*entities.User{}, nil
	}

	var users []*entities.User
	err := r.db.Connect().
		Where("id IN ?", userIDs).
		Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userRepositoryImpl) Update(tx *gorm.DB, updateData *entities.User) error {

	err := tx.
		Model(&entities.User{}).
		Where("id = ?", updateData.ID).
		Updates(updateData).Error
	if err != nil {
		return _exception.Handle("cannot update user", _exceptionType.Info)
	}
	return nil
}

func (r *userRepositoryImpl) Delete(tx *gorm.DB, userID uuid.UUID) error {

	err := tx.
		Model(&entities.User{}).
		Where("id = ?", userID).
		Update("is_deleted", true).Error
	if err != nil {
		return _exception.Handle("cannot delete user", _exceptionType.Info)
	}

	return nil
}
