package main

import (
	"cursor-response/handler"
	"cursor-response/service"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	service := service.NewPokemonService()
	handler := handler.NewHandler(service)

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/items-cursor", handler.GetPokemons)

	log.Fatal(http.ListenAndServe(":8080", r))
}
