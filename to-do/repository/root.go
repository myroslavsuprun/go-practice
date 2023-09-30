package repository

type Todo struct {
	Id        int
	Title     string
	Completed bool
	CreatedAt string
}

var todos []Todo
