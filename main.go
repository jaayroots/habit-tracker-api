package main

import (
	"github.com/jaayroots/go_base/config"
	"github.com/jaayroots/go_base/database"
	"github.com/jaayroots/go_base/server"
)

func main() {
	conf := config.ConfigGetting()
	db := database.NewPostgresDatabase(conf.Database)
	server := server.NewEchoServer(conf, db)

	server.Start()
}
