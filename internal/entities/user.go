package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/habit-tracker-api/enums"
)

type User struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey"`
	Email         string         `gorm:"type:varchar(128);unique;not null;"`
	Password      string         `gorm:"type:varchar(256);unique;not null;"`
	FirstName     string         `gorm:"type:varchar(128);not null;"`
	LastName      string         `gorm:"type:varchar(128);not null;"`
	UserType      enums.UserType `gorm:"type:varchar(20);default:user;not null"`
	IsCheckAttend bool           `gorm:"not null;default:false;"`
	IsDeleted     bool           `gorm:"not null;default:false;"`
	CreatedAt     time.Time      `gorm:"not null;autoCreateTime;"`
	UpdatedAt     time.Time      `gorm:"not null;autoUpdateTime;"`

	// HabitID     Habit       `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// CreatedByHabit []Habit `gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// UpdatedByHabit []Habit `gorm:"foreignKey:UpdatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// DeletedByHabit []Habit `gorm:"foreignKey:DeletedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	// CreatedByCheckin []Checkin `gorm:"foreignKey:CreatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// UpdatedByCheckin []Checkin `gorm:"foreignKey:UpdatedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	// DeletedByCheckin []Checkin `gorm:"foreignKey:DeletedBy;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

}

func (u User) GetID() uuid.UUID {
	return u.ID
}
