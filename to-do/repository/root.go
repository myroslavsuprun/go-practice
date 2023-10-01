package repository

import (
	"database/sql"
)

type Repository struct {
	db *sql.DB
}

type IRepository interface {
	Add(string) error
	Get(GetOpts) ([]Todo, error)
	Complete(int) (string, error)
	Remove(int) (string, error)
	Clear() error
}

func Get(db *sql.DB) IRepository {
	return &Repository{db: db}
}
