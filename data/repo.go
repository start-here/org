package data

import "log"

func NewFooRepo() FooRepo {
	return &fooRepo{}
}
func NewBarRepo() BarRepo {
	return &barRepo{}
}

type FooRepo interface {
	Get(int) (Foo, error)
}
type BarRepo interface {
	Get(int) (Bar, error)
}

type fooRepo struct{}

func (repo *fooRepo) Get(id int) (Foo, error) {
	log.Println("Get Foo")
	return Foo{Id: id}, nil
}

type barRepo struct{}

func (repo *barRepo) Get(id int) (Bar, error) {
	log.Println("Get Bar")
	return Bar{Id: id}, nil
}
