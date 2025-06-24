package main

import (
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/shahid-io/urban-waddle/cmd/api"
	"github.com/shahid-io/urban-waddle/config"
)

func main() {

	fmt.Println("urban waddle")

	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "urban-waddle-dev",
		AllowNativePasswords: true,
	}

	db, err := config.NewMySQLStore(cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB connected")

	server := api.NewAPIServer(":8080", db.DB)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
