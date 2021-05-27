package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/tikimcrzx/bookings/cmd/pkg/config"
	"github.com/tikimcrzx/bookings/cmd/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/test", handlers.Repo.Test)
	mux.Get("/single-room", handlers.Repo.SingleRoom)
	mux.Get("/deluxe-room", handlers.Repo.DeluxeRoom)
	mux.Get("/reservation", handlers.Repo.Reservation)
	mux.Get("/make-reservation", handlers.Repo.MakeReservation)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
