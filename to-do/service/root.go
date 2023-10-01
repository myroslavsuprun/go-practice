package service

import "to-do/repository"

type Service struct {
	r repository.IRepository
}

type IService interface {
	Add(string) error
	Clear() error
	Get() ([]TodoShow, error)
	GetToday() ([]TodoShow, error)
	Remove(int) (string, error)
	Complete(int) (string, error)
}

func New(r repository.IRepository) IService {
	return &Service{
		r: r,
	}
}
