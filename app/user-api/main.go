package main

import (
	"log"
	"os"

	"github.com/go-service-boilerplate/app/user-api/server"
	"github.com/go-service-boilerplate/config"
	db "github.com/go-service-boilerplate/infrastructure/db/mysql"
)

func main() {
	cfg, err := config.LoadConfig()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	dbClient, err := db.NewClient(cfg.DB)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	srv := server.New(dbClient)

	if err = srv.ListenAndServe(":8080"); err != nil {
		log.Println(err)
		os.Exit(1)
	}

}
