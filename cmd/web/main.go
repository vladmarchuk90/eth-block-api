package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/jellydator/ttlcache/v3"
	"github.com/vladmarchuk90/eth-block-api/pkg/config"
	"github.com/vladmarchuk90/eth-block-api/pkg/handlers"
	"github.com/vladmarchuk90/eth-block-api/pkg/models"
)

func main() {
	// initialize config
	configFilename, _ := filepath.Abs("config.json")
	app := config.NewConfig(configFilename)

	// using ttlcache as caching in-memory store, details at github.com/jellydator/ttlcache/v3
	cache := ttlcache.New[string, string]()
	app.Cache = cache

	// initialize handlers
	repo := handlers.NewRepo(app)
	handlers.NewHandlers(repo)

	// initialize models
	models.NewModels(app)

	// setup server
	srv := &http.Server{
		Handler: routes(app),
	}

	// run server
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
