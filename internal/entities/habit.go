package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/habit-tracker-api/enums"
	"gorm.io/gorm"
)

type Habit struct {
	ID          uint            `gorm:"primaryKey"`
	UserID      uuid.UUID       `gorm:"not null"`
	Title       string          `gorm:"type:varchar(255);not null"`
	Description string          `gorm:"type:text"`
	Frequency   enums.Frequency `gorm:"type:int;default:1;not null"`
	TargetCount uint            `gorm:"not null"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	CreatedBy uuid.UUID  `gorm:"not null"`
	UpdatedBy uuid.UUID  `gorm:"not null"`
	DeletedBy *uuid.UUID `gorm:"column:deleted_by"`

	Checkins []*Checkin `gorm:"foreignKey:HabitID;constraint:OnDelete:CASCADE"`
}

func (t *Habit) GetCreatedBy() uuid.UUID {
	return t.CreatedBy
}

func (t *Habit) GetUpdatedBy() uuid.UUID {
	return t.UpdatedBy
}

func (t *Habit) GetDeletedBy() uuid.UUID {
	if t.DeletedBy != nil {
		return *t.DeletedBy
	}
	return uuid.Nil
}

func (t *Habit) BeforeCreate(tx *gorm.DB) error {
	if err := setBlameableFieldsBeforeCreate(tx, &t.CreatedBy, &t.UpdatedBy); err != nil {
		return err
	}
	return nil
}

func (t *Habit) BeforeUpdate(tx *gorm.DB) error {
	if err := setBlameableFieldsBeforeUpdate(tx, &t.UpdatedBy); err != nil {
		return err
	}
	return nil
}

func (t *Habit) BeforeDelete(tx *gorm.DB) error {
	return setBlameableFieldsBeforeDelete(tx, t, t.ID)
}
