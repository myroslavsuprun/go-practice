package db

import (
	"os"
	"strconv"
)

// Default values
var (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "to-do"
)

func setDBConnectKeys() {
	if v, ok := os.LookupEnv("DB_HOST"); ok {
		host = v
	}
	if v, ok := os.LookupEnv("DB_PORT"); ok {
		port, _ = strconv.Atoi(v)
	}
	if v, ok := os.LookupEnv("DB_USER"); ok {
		user = v
	}
	if v, ok := os.LookupEnv("DB_PASSWORD"); ok {
		password = v
	}
	if v, ok := os.LookupEnv("DB_NAME"); ok {
		dbname = v
	}
}
