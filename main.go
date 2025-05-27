package main

import (
	"github.com/jaayroots/habit-tracker-api/config"
	"github.com/jaayroots/habit-tracker-api/database"
	"github.com/jaayroots/habit-tracker-api/server"
)

func main() {
	conf := config.ConfigGetting()
	db := database.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
