package repository

import (
	"errors"

	"github.com/google/uuid"
	databases "github.com/habit-tracker-api/database"
	"github.com/habit-tracker-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_exceptionType "github.com/habit-tracker-api/enums/exception"
	_exception "github.com/habit-tracker-api/exception"
)

type sessionRepositoryImpl struct {
	db     databases.Database
	logger echo.Logger
}

func NewSessionRepositoryImpl(db databases.Database, logger echo.Logger) SessionRepository {
	return &sessionRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

type SessionRepository interface {
	Create(tx *gorm.DB, session *entities.Session) error
	Delete(tx *gorm.DB, userID uuid.UUID) error
	FindByToken(token string) (*entities.Session, error)
}

// func (r *sessionRepositoryImpl) Create(session *entities.Session) (*entities.Session, error) {

// 	err := r.db.Connect().
// 		Create(session).Error

//		if err != nil {
//			return nil, _authException.CannotCreateSession()
//		}
//		return session, nil
//	}
func (r *sessionRepositoryImpl) Create(tx *gorm.DB, session *entities.Session) error {
	err := tx.Create(session).Error
	if err != nil {
		return _exception.Handle("cannot create session", _exceptionType.Info)
	}
	return err
}

func (r *sessionRepositoryImpl) Delete(tx *gorm.DB, userID uuid.UUID) error {

	err := tx.Delete(&entities.Session{}, "user_id = ?", userID).Error

	if err != nil {
		return _exception.Handle("cannot create session", _exceptionType.Info)
	}

	return nil
}

func (r *sessionRepositoryImpl) FindByToken(token string) (*entities.Session, error) {

	session := new(entities.Session)

	err := r.db.Connect().
		Where("token = ?", token).
		First(session).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return session, nil
}
