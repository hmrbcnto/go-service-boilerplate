package main

import (
	"log"
	"os"

	"github.com/go-service-boilerplate/app/user-api/server"
	db "github.com/go-service-boilerplate/infrastructure/db/mysql"
)

func main() {
	dbClient, err := db.NewClient()
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
