package main

import (
	"api-golang-test/configs"
	"api-golang-test/handlers"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Post("/", handlers.Create)
	r.Get("/", handlers.List)
	r.Get("/{id}", handlers.Find)
	r.Put("/{id}", handlers.Update)
	r.Delete("/{id}", handlers.Delete)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), r)
}
