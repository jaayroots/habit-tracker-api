package migration

import (
	"github.com/jaayroots/habit-tracker-api/entities"
	"gorm.io/gorm"
)

// Up_20250712172905_add_user_table applies the migration.
func Up_20250712172905_add_user_table(tx *gorm.DB) error {
	tx.Migrator().CreateTable(&entities.User{})
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	return nil
}

// Down_20250712172905_add_user_table rolls back the migration.
func Down_20250712172905_add_user_table(tx *gorm.DB) error {
	tx.Migrator().DropTable(&entities.User{})
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	return nil
}
