package migration

import (
	"github.com/habit-tracker-api/entities"
	"gorm.io/gorm"
)

// Up_20250712172922_add_session_table applies the migration.
func Up_20250712172922_add_session_table(tx *gorm.DB) error {
	tx.Migrator().CreateTable(&entities.Session{})
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	return nil
}

// Down_20250712172922_add_session_table rolls back the migration.
func Down_20250712172922_add_session_table(tx *gorm.DB) error {
	tx.Migrator().DropTable(&entities.Session{})
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	return nil
}
