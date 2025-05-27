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

	HabitID Habit `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedByHabit []Habit `gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UpdatedByHabit []Habit `gorm:"foreignKey:UpdatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DeletedByHabit []Habit `gorm:"foreignKey:DeletedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedByCheckIn []CheckIn `gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	UpdatedByCheckIn []CheckIn `gorm:"foreignKey:UpdatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	DeletedByCheckIn []CheckIn `gorm:"foreignKey:DeletedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (u User) GetID() uint {
	return u.ID
}
