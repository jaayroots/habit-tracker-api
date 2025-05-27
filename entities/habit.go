package entities

import (
	"time"

	"github.com/jaayroots/habit-tracker-api/enums"
	"gorm.io/gorm"
)

type Habit struct {
	ID          uint            `gorm:"primaryKey"`
	UserID      uint            `gorm:"not null"`
	Title       string          `gorm:"type:varchar(255);not null"`
	Description string          `gorm:"type:text"`
	Frequency   enums.Frequency `gorm:"type:int;default:1;not null"`
	TargetCount uint            `gorm:"not null"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	CreatedBy uint  `gorm:"not null"`
	UpdatedBy uint  `gorm:"not null"`
	DeletedBy *uint `gorm:"column:deleted_by"`

	HabitID Checkin `gorm:"foreignKey:HabitID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (t *Habit) GetCreatedBy() uint {
	return t.CreatedBy
}

func (t *Habit) GetUpdatedBy() uint {
	return t.UpdatedBy
}

func (t *Habit) GetDeletedBy() uint {
	if t.DeletedBy != nil {
		return *t.DeletedBy
	}
	return 0
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
