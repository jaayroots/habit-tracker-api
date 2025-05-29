package main

import (
	"github.com/jaayroots/habit-tracker-api/config"
	"github.com/jaayroots/habit-tracker-api/database"
	"github.com/jaayroots/habit-tracker-api/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := database.NewPostgresDatabase(conf.Database)

	tx := db.Connect().Begin()

	userMigration(tx)
	sessionMigration(tx)
	habitMigration(tx)
	CheckinMigration(tx)

	tx.Commit()
	if tx.Error != nil {
		tx.Rollback()
		panic(tx.Error)
	}
}

func userMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.User{})
}

func sessionMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Session{})
}

func habitMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Habit{})
}

func CheckinMigration(tx *gorm.DB) {
	tx.Migrator().CreateTable(&entities.Checkin{})
}
