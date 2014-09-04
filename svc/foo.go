package svc

import (
	"localhost/scratch/org/data"
	"log"
)

func NewFooService(repo data.FooRepo) FooService {
	return fooService{
		repo: repo,
	}
}

type FooService interface {
	Foo(int) (data.Foo, error)
}
type fooService struct {
	repo data.FooRepo
}

func (s fooService) Foo(id int) (data.Foo, error) {
	foo, err := s.repo.Get(id)
	log.Println("Foo()", foo)
	return foo, err
}
