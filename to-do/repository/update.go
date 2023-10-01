package repository

import (
	"to-do/db"
)

const updateSql = `
		UPDATE todo
		SET completed = $1
		WHERE id = $2
		RETURNING title;
`

func (r *Repository) Complete(id int) (string, error) {
	var title string
	err := db.DB.QueryRow(updateSql, true, id).Scan(&title)
	if err != nil {
		return title, err
	}

	return title, nil
}
