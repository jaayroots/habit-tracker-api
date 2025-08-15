package packImport

import (
	"github.com/jaayroots/habit-tracker-api/command/migration/list"
	"gorm.io/gorm"
)

var migrationUpFuncs = map[string]func(*gorm.DB) error{
	"20250712172905_add_user_table": migration.Up_20250712172905_add_user_table,
	"20250712172922_add_session_table": migration.Up_20250712172922_add_session_table,
	"20250712172940_add_habit_table": migration.Up_20250712172940_add_habit_table,
	"20250712173007_add_checkin_table": migration.Up_20250712173007_add_checkin_table,
	"20250814142845_add_table_user_contact": migration.Up_20250814142845_add_table_user_contact,

}

var migrationDownFuncs = map[string]func(*gorm.DB) error{
	"20250712172905_add_user_table": migration.Down_20250712172905_add_user_table,
	"20250712172922_add_session_table": migration.Down_20250712172922_add_session_table,
	"20250712172940_add_habit_table": migration.Down_20250712172940_add_habit_table,
	"20250712173007_add_checkin_table": migration.Down_20250712173007_add_checkin_table,
	"20250814142845_add_table_user_contact": migration.Down_20250814142845_add_table_user_contact,

}

func MappingUpFuncMigration() map[string]func(*gorm.DB) error {
	return migrationUpFuncs
}

func MappingDownFuncMigration() map[string]func(*gorm.DB) error {
	return migrationDownFuncs
}
