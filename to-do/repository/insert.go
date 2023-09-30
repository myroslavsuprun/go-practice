package repository

import (
	"to-do/db"
)

const insertSql = `
		INSERT INTO todo (title, completed)
		VALUES ($1, $2)
`

func Add(title string) error {
	_, err := db.DB.Exec(insertSql, title, false)
	if err != nil {
		return err
	}

	return nil
}
