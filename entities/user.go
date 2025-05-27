package entities

import (
	"time"
)

type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	Email     string    `gorm:"type:varchar(128);unique;not null;"`
	Password  string    `gorm:"type:varchar(256);unique;not null;"`
	FirstName string    `gorm:"type:varchar(128);not null;"`
	LastName  string    `gorm:"type:varchar(128);not null;"`
	IsDeleted bool      `gorm:"not null;default:false;"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime;"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime;"`
}

func (u User) GetID() uint {
	return u.ID
}
