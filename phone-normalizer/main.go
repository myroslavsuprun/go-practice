package main

import (
	"fmt"
	"log"
	dbService "phone-normalizer/db"
	"phone-normalizer/normalizer"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	db, err := dbService.Connect()
	if err != nil {
		log.Fatal("Error connecting to the database.")
	}
	defer dbService.Disconnect(db)

	res, err := db.Query(`SELECT id, phone_number FROM phone_numbers;`)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()

	numbers := make(map[string]string)

	for res.Next() {
		var id string
		var phoneNumber string
		err = res.Scan(&id, &phoneNumber)
		if err != nil {
			log.Fatal(err)
		}

		numbers[id] = normalizer.Normalize(phoneNumber)
	}

	for id, phoneNumber := range numbers {
		_, err := db.Exec("UPDATE phone_numbers SET phone_number = $1 WHERE id = $2", phoneNumber, id)
		if err != nil {
			log.Fatal(err)
		}
	}

	_, err = db.Exec(`
		DELETE FROM phone_numbers
		WHERE id IN (
			SELECT id
			FROM (
				SELECT id,
					ROW_NUMBER() OVER (
						PARTITION BY phone_number
						ORDER BY id
					) AS row_num
				FROM phone_numbers
			) t
			WHERE t.row_num > 1
		);
	`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done!")
}
