package service

import "to-do/repository"

func (s *Service) Get() ([]TodoShow, error) {
	todos, err := s.r.Get(repository.GetOpts{
		Completed: false,
	})
	if err != nil {
		return nil, err
	}

	todosShow := getTodoShow(todos)

	return todosShow, nil

}

type TodoShow struct {
	Title string
	Id    int
}

func getTodoShow(todos []repository.Todo) []TodoShow {
	var todosShow []TodoShow
	for _, v := range todos {
		todosShow = append(todosShow, TodoShow{Title: v.Title, Id: v.Id})
	}

	return todosShow
}