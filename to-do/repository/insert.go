package repository

const insertSql = `
		INSERT INTO todo (title, completed)
		VALUES ($1, $2)
`

func (r *Repository) Add(title string) error {
	_, err := r.db.Exec(insertSql, title, false)
	if err != nil {
		return err
	}

	return nil
}
