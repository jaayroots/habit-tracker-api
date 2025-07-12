package packImport

import (
	
	"gorm.io/gorm"
)

var migrationUpFuncs = map[string]func(*gorm.DB) error{

}

var migrationDownFuncs = map[string]func(*gorm.DB) error{

}

func MappingUpFuncMigration() map[string]func(*gorm.DB) error {
	return migrationUpFuncs
}

func MappingDownFuncMigration() map[string]func(*gorm.DB) error {
	return migrationDownFuncs
}
