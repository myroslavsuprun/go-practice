package service

import "to-do/repository"

var complete Complete = func(id int) (string, error) {
	title, err := repository.Completed(id)
	if err != nil {
		return title, err
	}

	return title, nil
}
