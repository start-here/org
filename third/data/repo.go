package data

import "log"

func NewBazRepo() BazRepo {
	return bazRepo{}
}

type BazRepo interface {
	Get(int) (Baz, error)
}
type bazRepo struct{}

func (s bazRepo) Get(id int) (Baz, error) {
	log.Println("Get Baz")
	return Baz{Id: id}, nil
}
