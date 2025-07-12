package migration

import (
	"github.com/jaayroots/habit-tracker-api/entities"
	"gorm.io/gorm"
)

// Up_20250712173007_add_checkin_table applies the migration.
func Up_20250712173007_add_checkin_table(tx *gorm.DB) error {
	tx.Migrator().CreateTable(&entities.Checkin{})
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	return nil
}

// Down_20250712173007_add_checkin_table rolls back the migration.
func Down_20250712173007_add_checkin_table(tx *gorm.DB) error {
	tx.Migrator().DropTable(&entities.Checkin{})
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	return nil
}
