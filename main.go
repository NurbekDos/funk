package main

import (
	"log"
	"time"

	"github.com/NurbekDos/funk/internal/db"
	"github.com/NurbekDos/funk/internal/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf(".env load error: %s\n" + err.Error())
	}

	err = db.RunMigrations()
	if err != nil {
		time.Sleep(time.Second * 10)

		err = db.RunMigrations()
		if err != nil {
			log.Printf("RunMigrations error: %s\n", err.Error())
			return
		}
	}

	err = db.ConnectToPostgres()
	if err != nil {
		log.Printf("OpenConnection error: %s\n", err.Error())
		return
	}
	defer db.Close()

	server.Engine()
}
