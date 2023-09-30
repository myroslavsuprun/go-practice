package repository

import (
	"to-do/db"
)

const removeSql = `
	DELETE FROM todo
	WHERE id = $1
	RETURNING title;
`

func Remove(id int) (string, error) {
	var title string
	err := db.DB.QueryRow(removeSql, id).Scan(&title)
	if err != nil {
		return title, err
	}

	return title, nil
}
