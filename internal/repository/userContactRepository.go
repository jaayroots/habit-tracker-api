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

type userContactRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewUserContactRepository(db databases.Database, logger echo.Logger) UserContactRepository {
	return &userContactRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

type UserContactRepository interface {
	FindByID(userID uuid.UUID) (*entities.UserContact, error)
	FindByIDs(userIDs []uuid.UUID) ([]*entities.User, error)
	FindByUserID(userID uuid.UUID) ([]*entities.UserContact, error)
	Create(tx *gorm.DB, userContact []entities.UserContact) error
	Update(tx *gorm.DB, userContact []entities.UserContact) error
	Delete(tx *gorm.DB, userContact []entities.UserContact) error
}

func (r *userContactRepositoryImpl) FindByID(userContactID uuid.UUID) (*entities.UserContact, error) {

	userContact := new(entities.UserContact)

	err := r.db.Connect().
		Where("id = ?", userContactID).
		First(userContact).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return userContact, nil
}

func (r *userContactRepositoryImpl) FindByIDs(userIDs []uuid.UUID) ([]*entities.User, error) {

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

func (r *userContactRepositoryImpl) FindByUserID(userID uuid.UUID) ([]*entities.UserContact, error) {
	var userContacts []*entities.UserContact

	err := r.db.Connect().
		Where("user_id = ?", userID).
		Find(&userContacts).Error
	if err != nil {
		return nil, err
	}

	return userContacts, nil
}

func (r *userContactRepositoryImpl) Create(tx *gorm.DB, userContact []entities.UserContact) error {
	if len(userContact) == 0 {
		return nil
	}

	err := tx.Create(userContact).Error
	if err != nil {
		return _exception.Handle("cannot create user contact", _exceptionType.Info)
	}
	return nil
}

func (r *userContactRepositoryImpl) Update(tx *gorm.DB, userContact []entities.UserContact) error {

	if len(userContact) == 0 {
		return nil
	}

	for _, contact := range userContact {
		if err := tx.Model(&entities.UserContact{}).
			Where("id = ?", contact.ID).
			Updates(contact).Error; err != nil {
			return err
		}
	}
	return nil
}

func (r *userContactRepositoryImpl) Delete(tx *gorm.DB, userContact []entities.UserContact) error {
	if len(userContact) == 0 {
		return nil
	}
	for _, contact := range userContact {
		if err := tx.
			Where("id = ?", contact.ID).
			Where("user_id = ?", contact.UserID).
			Delete(&entities.UserContact{}).
			Error; err != nil {
			return err
		}
	}
	return nil
}
