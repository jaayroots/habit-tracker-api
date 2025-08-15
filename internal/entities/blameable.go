package entities

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/jaayroots/habit-tracker-api/pctxkeys"
)

func setBlameableFieldsBeforeCreate(tx *gorm.DB, createdBy *uuid.UUID, updatedBy *uuid.UUID) error {
	userID, ok := tx.Statement.Context.Value(pctxkeys.ContextKeyUserID).(uuid.UUID)
	if !ok {
		return errors.New("userID not found in context")
	}
	*createdBy = userID
	*updatedBy = userID
	return nil
}

func setBlameableFieldsBeforeUpdate(tx *gorm.DB, updatedBy *uuid.UUID) error {
	userID, ok := tx.Statement.Context.Value(pctxkeys.ContextKeyUserID).(uuid.UUID)
	if !ok {
		return errors.New("userID not found in context")
	}
	*updatedBy = userID
	return nil
}

func setBlameableFieldsBeforeDelete(tx *gorm.DB, model interface{}, id uint) error {
	userID, ok := tx.Statement.Context.Value(pctxkeys.ContextKeyUserID).(uuid.UUID)
	if !ok {
		return errors.New("userID not found in context")
	}
	err := tx.Model(model).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_by": userID,
			"updated_at": time.Now(),
		}).Error
	return err
}
