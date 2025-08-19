package migration

import (
	"github.com/habit-tracker-api/entities"
	"gorm.io/gorm"
)

// Up_20250712172940_add_habit_table applies the migration.
func Up_20250712172940_add_habit_table(tx *gorm.DB) error {
	tx.Migrator().CreateTable(&entities.Habit{})
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	return nil
}

// Down_20250712172940_add_habit_table rolls back the migration.
func Down_20250712172940_add_habit_table(tx *gorm.DB) error {
	tx.Migrator().DropTable(&entities.Habit{})
	if tx.Error != nil {
		tx.Rollback()
		return tx.Error
	}
	return nil
}
