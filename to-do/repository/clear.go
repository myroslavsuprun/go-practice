package repository

const clearSql = `
		DELETE FROM todo;
		ALTER SEQUENCE todo_id_seq RESTART WITH 1;
`

func (r *Repository) Clear() error {
	_, err := r.db.Exec(clearSql)
	if err != nil {
		return err
	}

	return nil

}
