package repository

import "github.com/jaayroots/habit-tracker-api/entities"

type SessionRepository interface {
	Create(session *entities.Session) (*entities.Session, error)
	Delete(userID uint) error
	FindByToken(token string) (*entities.Session, error)
}
