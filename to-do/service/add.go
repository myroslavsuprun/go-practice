package service

import "to-do/repository"

func add(title string) error {
	err := repository.Add(title)
	if err != nil {
		return err
	}
	return nil
}
