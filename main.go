package main

import (
	"fmt"
	"localhost/scratch/org/api"
	"localhost/scratch/org/data"
	"localhost/scratch/org/svc"
	thrdata "localhost/scratch/org/third/data"
	thrsvc "localhost/scratch/org/third/svc"
	"net/http"
	"time"
)

func main() {

	bazRepo := thrdata.NewBazRepo()
	bazSvc := thrsvc.NewBazService(bazRepo)
	fooRepo := data.NewFooRepo()
	fooSvc := svc.NewFooService(fooRepo)
	barRepo := data.NewBarRepo()
	barSvc := svc.NewBarService(barRepo, fooSvc, bazSvc)

	http.HandleFunc("/foo", api.FooHandler)
	http.HandleFunc("/bar", api.BarHandler)
	http.HandleFunc("/baz", api.BazHandler)

	fmt.Println(barSvc.Bar(time.Now().Nanosecond()))
}
