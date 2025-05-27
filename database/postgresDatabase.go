package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/jaayroots/go_base/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgesDatabase struct {
	*gorm.DB
}

var (
	postgesDatabaseInstance *postgesDatabase
	once                    sync.Once
)

func NewPostgresDatabase(conf *config.Database) Database {
	once.Do(func() {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s search_path=%s",
			conf.Host,
			conf.User,
			conf.Password,
			conf.DBName,
			conf.Port,
			conf.SSLMode,
			conf.Schema,
		)

		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err != nil {
			panic(err)
		}

		log.Printf("Connected to database %s", conf.DBName)

		postgesDatabaseInstance = &postgesDatabase{conn}
	})
	return postgesDatabaseInstance
}

func (db *postgesDatabase) Connect() *gorm.DB {
	return postgesDatabaseInstance.DB
}
