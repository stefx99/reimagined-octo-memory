package controller

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func Router() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Handle("/img/*", http.FileServer(http.Dir("static")))

	r.Get("/", homeHandler)
	r.Get("/reservation", reservationHandler)
	r.Get("/reservation/success", reservationSuccessHandler)
	r.Post("/reservation/submit", submitReservation)
	//r.Get("/profile/reservation", profileReservationHandler)
	r.Get("/login", loginHandler)
	r.Post("/login", loginAuthHandler)

	return r
}
