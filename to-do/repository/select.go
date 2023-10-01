package repository

import (
	"time"
)

type GetOpts struct {
	Completed bool
	From      string
}

type Todo struct {
	Id        int
	Title     string
	Completed bool
	CreatedAt string
}

const selectSql = `
		SELECT id, title, completed, created_at
		FROM todo
		WHERE completed = $1
		AND created_at > $2
`

func (r *Repository) Get(opts GetOpts) ([]Todo, error) {
	if opts.From == "" {
		opts.From = time.RFC3339
	}

	rows, err := r.db.Query(selectSql, opts.Completed, opts.From)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var todo Todo
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Completed, &todo.CreatedAt)
		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil
}
