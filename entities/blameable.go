package entities

import (
	"database/sql"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/jaayroots/habit-tracker-api/pctxkeys"
)

func setBlameableFieldsBeforeCreate(tx *gorm.DB, createdBy *uint, updatedBy *uint) error {
	userID, ok := tx.Statement.Context.Value(pctxkeys.ContextKeyUserID).(uint)
	if !ok {
		return errors.New("userID not found in context")
	}
	*createdBy = userID
	*updatedBy = userID
	return nil
}

func setBlameableFieldsBeforeUpdate(tx *gorm.DB, updatedBy *uint) error {
	userID, ok := tx.Statement.Context.Value(pctxkeys.ContextKeyUserID).(uint)
	if !ok {
		return errors.New("userID not found in context")
	}
	*updatedBy = userID
	return nil
}

func setBlameableFieldsBeforeDelete(tx *gorm.DB, model interface{}, id uint) error {
	userID, ok := tx.Statement.Context.Value(pctxkeys.ContextKeyUserID).(uint)
	if !ok {
		return errors.New("userID not found in context")
	}
	err := tx.Model(model).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"deleted_by": sql.NullInt64{Int64: int64(userID), Valid: true},
			"updated_at": time.Now(),
		}).Error
	return err
}
