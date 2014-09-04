package svc

import (
	"localhost/scratch/org/third/data"
	"log"
)

func NewBazService(repo data.BazRepo) BazService {
	return bazService{
		repo: repo,
	}
}

type BazService interface {
	Baz(int) (data.Baz, error)
}
type bazService struct {
	repo data.BazRepo
}

func (s bazService) Baz(id int) (data.Baz, error) {
	baz, err := s.repo.Get(id)
	log.Println("Baz()", baz)
	return baz, err
}
