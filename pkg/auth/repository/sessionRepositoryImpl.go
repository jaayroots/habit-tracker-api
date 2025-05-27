package repository

import (
	"errors"

	databases "github.com/jaayroots/habit-tracker-api/database"
	"github.com/jaayroots/habit-tracker-api/entities"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_authException "github.com/jaayroots/habit-tracker-api/pkg/auth/exception"
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

func (r *sessionRepositoryImpl) Create(session *entities.Session) (*entities.Session, error) {

	err := r.db.Connect().
		Create(session).Error

	if err != nil {
		return nil, _authException.CannotCreateSession()
	}
	return session, nil
}

func (r *sessionRepositoryImpl) Delete(userID uint) error {

	err := r.db.Connect().
		Delete(&entities.Session{}, "user_id = ?", userID).Error

	if err != nil {
		return _authException.CannotCreateSession()
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
