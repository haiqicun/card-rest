package handler

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/haiqicun/card-rest/pkg/repository"
)

var dbInstance repository.Database

//NewHandler creates a new instance of "Handler" and map to the Routes.
func NewHandler(db repository.Database) http.Handler {
	router := chi.NewRouter()
	dbInstance = db
	router.MethodNotAllowed(methodNotAllowedHandler)
	router.NotFound(notFoundHandler)
	router.Route("/deck", deck)

	return router
}

func methodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(405)
	render.Render(w, r, ErrMethodNotAllowed)
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(400)
	render.Render(w, r, ErrNotFound)
}
