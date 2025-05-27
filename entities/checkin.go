package entities

import (
	"time"

	"gorm.io/gorm"
)

type Checkin struct {
	ID      uint `gorm:"primaryKey"`
	HabitID uint `gorm:"not null"`
	Date    *time.Time

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	CreatedBy uint  `gorm:"not null"`
	UpdatedBy uint  `gorm:"not null"`
	DeletedBy *uint `gorm:"column:deleted_by"`
}

func (t *Checkin) GetCreatedBy() uint {
	return t.CreatedBy
}

func (t *Checkin) GetUpdatedBy() uint {
	return t.UpdatedBy
}

func (t *Checkin) GetDeletedBy() uint {
	if t.DeletedBy != nil {
		return *t.DeletedBy
	}
	return 0
}

func (t *Checkin) BeforeCreate(tx *gorm.DB) error {
	if err := setBlameableFieldsBeforeCreate(tx, &t.CreatedBy, &t.UpdatedBy); err != nil {
		return err
	}
	return nil
}

func (t *Checkin) BeforeUpdate(tx *gorm.DB) error {
	if err := setBlameableFieldsBeforeUpdate(tx, &t.UpdatedBy); err != nil {
		return err
	}
	return nil
}

func (t *Checkin) BeforeDelete(tx *gorm.DB) error {
	return setBlameableFieldsBeforeDelete(tx, t, t.ID)
}
