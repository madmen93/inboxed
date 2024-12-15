package server

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/go-chi/chi/v5"
	"github.com/madmen93/inboxed/cmd/server/handlers"
)

func SetUpRoutes() http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(CORSMiddleware)
	router.Get("/health", handlers.HealthCheckHandler)
	router.Get("/emails", handlers.GetEmailsFiltered)
	router.Get("/emails/{id}", handlers.GetEmailByID)
	return router
}
