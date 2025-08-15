package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/jaayroots/habit-tracker-api/enums"
	"gorm.io/gorm"
)

type UserContact struct {
	ID          uuid.UUID         `gorm:"type:uuid;primaryKey"`
	UserID      uuid.UUID         `gorm:"not null"`
	ContactType enums.ContactType `gorm:"type:varchar(20);not null"`
	Value       string            `gorm:"type:varchar(128);not null;"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`

	// CreatedBy uuid.UUID  `gorm:"not null"`
	// UpdatedBy uuid.UUID  `gorm:"not null"`
	// DeletedBy *uuid.UUID `gorm:"column:deleted_by"`
}

// func (t *Checkin) GetCreatedBy() uuid.UUID {
// 	return t.CreatedBy
// }

// func (t *Checkin) GetUpdatedBy() uuid.UUID {
// 	return t.UpdatedBy
// }

// func (t *Checkin) GetDeletedBy() uuid.UUID {
// 	if t.DeletedBy != nil {
// 		return *t.DeletedBy
// 	}
// 	return uuid.Nil
// }

// func (t *Checkin) BeforeCreate(tx *gorm.DB) error {
// 	if err := setBlameableFieldsBeforeCreate(tx, &t.CreatedBy, &t.UpdatedBy); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (t *Checkin) BeforeUpdate(tx *gorm.DB) error {
// 	if err := setBlameableFieldsBeforeUpdate(tx, &t.UpdatedBy); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (t *Checkin) BeforeDelete(tx *gorm.DB) error {
// 	return setBlameableFieldsBeforeDelete(tx, t, t.ID)
// }
