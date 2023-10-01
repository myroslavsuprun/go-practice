package repository

const removeSql = `
	DELETE FROM todo
	WHERE id = $1
	RETURNING title;
`

func (r *Repository) Remove(id int) (string, error) {
	var title string
	err := r.db.QueryRow(removeSql, id).Scan(&title)
	if err != nil {
		return title, err
	}

	return title, nil
}
