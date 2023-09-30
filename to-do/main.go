package main

import (
	"log"

	"to-do/cmd"
	dbC "to-do/db"
	"to-do/service"

	"github.com/joho/godotenv"
)

func main() {
	db, err := dbC.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer dbC.Disconnect(db)

	actions := service.New()

	err = cmd.Execute(actions)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	godotenv.Load()
}
