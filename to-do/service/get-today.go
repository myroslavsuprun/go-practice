package service

import (
	"time"
	"to-do/repository"
)

func (s *Service) GetToday() ([]TodoShow, error) {
	todos, err := s.r.Get(repository.GetOpts{
		Completed: true,
		From:      fromToday(),
	})
	if err != nil {
		return nil, err
	}

	todosShow := getTodoShow(todos)

	return todosShow, nil
}

func fromToday() string {
	currT := time.Now()

	t := time.Date(currT.Year(), currT.Month(), currT.Day(), 0, 0, 0, 0, currT.Location())

	return t.Format(time.RFC3339)

}
