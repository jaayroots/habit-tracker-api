package repository

import "github.com/jaayroots/habit-tracker-api/entities"

type UserRepository interface {
	Create(user *entities.User) (*entities.User, error)
	FindByEmail(email string) (*entities.User, error)
	FindByID(userID uint) (*entities.User, error)
	FindByIDs(userIDs []uint) ([]*entities.User, error)
	Update(userID uint, user *entities.User) (*entities.User, error)
	Delete(userID uint) error
}
