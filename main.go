package main

import (
	"fmt"
	"net/http"

	"github.com/Babiel09/configs"
	"github.com/Babiel09/handlers"
	"github.com/go-chi/chi"
)

func main() {
	err := configs.Load()
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	r.Get("/people", handlers.ListAll)
	r.Get("/people/{id}", handlers.Read)
	r.Post("/people", handlers.Create)
	r.Put("/people/{id}", handlers.UpdateInformations)
	r.Delete("/people/{id}", handlers.Delte)

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetAPIPort()), r)
}
