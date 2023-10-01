package main

import (
	"log"

	"to-do/cmd"
	dbC "to-do/db"
	"to-do/repository"
	"to-do/service"

	"github.com/joho/godotenv"
)

func main() {
	db, err := dbC.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer dbC.Disconnect(db)

	r := repository.Get(db)
	s := service.New(r)

	err = cmd.Execute(s)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	godotenv.Load()
}
