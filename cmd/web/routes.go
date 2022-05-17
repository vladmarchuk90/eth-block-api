package main

import (
	"net/http"
	"time"

	"github.com/vladmarchuk90/eth-block-api/pkg/config"
	"github.com/vladmarchuk90/eth-block-api/pkg/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// routes - setup handlers for routes and used middleware
func routes(app *config.AppConfig) http.Handler {

	// chi is lightweight router, details at https://github.com/go-chi/chi
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Timeout(60 * time.Second))

	mux.Get("/api/block/{blockNumber}/total", handlers.Repo.GetBlockInfo)

	return mux
}
