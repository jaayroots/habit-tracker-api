package main

import (
	"github.com/jaayroots/go_base/config"
	"github.com/jaayroots/go_base/database"
	"github.com/jaayroots/go_base/entities"
	"gorm.io/gorm"
)

func main() {
	conf := config.ConfigGetting()
	db := database.NewPostgresDatabase(conf.Database)

	tx := db.Connect().Begin()

	userMigration(tx)
	sessionMigration(tx)

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
