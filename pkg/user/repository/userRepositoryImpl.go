package repository

import (
	"errors"

	databases "github.com/jaayroots/habit-tracker-api/database"
	"github.com/jaayroots/habit-tracker-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_userException "github.com/jaayroots/habit-tracker-api/pkg/user/exception"
)

type userRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewUserRepositoryImpl(db databases.Database, logger echo.Logger) UserRepository {
	return &userRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

func (r *userRepositoryImpl) Create(user *entities.User) (*entities.User, error) {

	err := r.db.Connect().
		Create(user).Error

	if err != nil {
		return nil, _userException.CannotCreateUser()
	}
	return user, nil
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

func (r *userRepositoryImpl) FindByID(userID uint) (*entities.User, error) {

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

func (r *userRepositoryImpl) FindByIDs(userIDs []uint) ([]*entities.User, error) {

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

func (r *userRepositoryImpl) Update(userID uint, updateData *entities.User) (*entities.User, error) {

	user := new(entities.User)

	err := r.db.Connect().
		Model(&entities.User{}).
		Where("id = ?", userID).
		Updates(user).Error
	if err != nil {
		return nil, _userException.CannotUpdateUser()
	}

	return user, nil
}

func (r *userRepositoryImpl) Delete(userID uint) error {

	err := r.db.Connect().
		Model(&entities.User{}).
		Where("id = ?", userID).
		Update("is_deleted", true).Error
	if err != nil {
		return _userException.CannotDeleteUser()
	}

	return nil
}
