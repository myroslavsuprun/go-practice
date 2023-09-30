package repository

import (
	"to-do/db"
)

const clearSql = `
		DELETE FROM todo;
		ALTER SEQUENCE todo_id_seq RESTART WITH 1;
`

func Clear() error {
	_, err := db.DB.Exec(clearSql)
	if err != nil {
		return err
	}

	return nil
}
