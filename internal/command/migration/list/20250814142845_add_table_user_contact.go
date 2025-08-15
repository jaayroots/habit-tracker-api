package migration

import (
	"github.com/jaayroots/habit-tracker-api/entities"
	"gorm.io/gorm"
)

// Up_20250814142845_add_table_user_contact applies the migration.
func Up_20250814142845_add_table_user_contact(tx *gorm.DB) error {
	tx.Migrator().CreateTable(&entities.UserContact{})
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	return nil
}

// Down_20250814142845_add_table_user_contact rolls back the migration.
func Down_20250814142845_add_table_user_contact(tx *gorm.DB) error {
	tx.Migrator().DropTable(&entities.UserContact{})
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	return nil
}
