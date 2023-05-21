package handlers

import (
	"github.com/go-chi/chi/v5"
	"go-test-inhousead/internal/config"
	"go-test-inhousead/internal/service"
)

type App struct {
	*chi.Mux
	Config config.Config
}

func NewApp(cfg config.Config) *App {

	app := &App{
		Mux:    chi.NewRouter(),
		Config: cfg,
	}

	app.Route("/", func(r chi.Router) {

		r.Get("/", service.PingHandler)
		r.Get("/{id}", SearchHandler)
		r.Get("/low", LowHandler)
		r.Get("/slow", SlowHandler)
		r.Get("/admin", AdminHandler)
	})

	return app
}
