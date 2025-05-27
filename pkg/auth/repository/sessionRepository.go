package repository

import "github.com/jaayroots/go_base/entities"

type SessionRepository interface {
	Create(session *entities.Session) (*entities.Session, error)
	Delete(userID uint) error
	FindByToken(token string) (*entities.Session, error)
}
