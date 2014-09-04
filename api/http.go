package api

import (
	"localhost/scratch/org/data"
	"localhost/scratch/org/svc"
	thrdata "localhost/scratch/org/third/data"
	thrsvc "localhost/scratch/org/third/svc"
	"net/http"
)

var (
	bazRepo thrdata.BazRepo
	bazSvc  thrsvc.BazService
	fooRepo data.FooRepo
	fooSvc  svc.FooService
	barRepo data.BarRepo
	barSvc  svc.BarService
)

func init() {
	bazRepo = thrdata.NewBazRepo()
	bazSvc = thrsvc.NewBazService(bazRepo)
	fooRepo = data.NewFooRepo()
	fooSvc = svc.NewFooService(fooRepo)
	barRepo = data.NewBarRepo()
	barSvc = svc.NewBarService(barRepo, fooSvc, bazSvc)
}

func FooHandler(w http.ResponseWriter, req *http.Request) {
	fooSvc.Foo(1)
}
func BarHandler(w http.ResponseWriter, req *http.Request) {
	barSvc.Bar(2)
}
func BazHandler(w http.ResponseWriter, req *http.Request) {
	bazSvc.Baz(3)
}
