package svc

import (
	"localhost/scratch/org/data"
	baz "localhost/scratch/org/third/svc"
	"log"
)

func NewBarService(repo data.BarRepo, fooSvc FooService, bazSvc baz.BazService) BarService {
	return barService{
		repo:   repo,
		fooSvc: fooSvc,
		bazSvc: bazSvc,
	}
}

type BarService interface {
	Bar(int) (data.Bar, error)
}
type barService struct {
	repo   data.BarRepo
	fooSvc FooService
	bazSvc baz.BazService
}

func (s barService) Bar(id int) (data.Bar, error) {
	_, err := s.fooSvc.Foo(id)
	bar, err := s.repo.Get(id)
	log.Println("Bar()", bar)
	_, err = s.bazSvc.Baz(id)
	return bar, err
}
