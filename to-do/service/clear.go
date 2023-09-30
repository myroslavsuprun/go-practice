package service

import "to-do/repository"

func clear() error {
	err := repository.Clear()
	if err != nil {
		return err
	}
	return nil
}
