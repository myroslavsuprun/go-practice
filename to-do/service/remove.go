package service

import "to-do/repository"

func remove(id int) (string, error) {
	title, err := repository.Remove(id)
	if err != nil {
		return title, err
	}

	return title, nil

}
