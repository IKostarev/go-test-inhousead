package main

import (
	"go-test-inhousead/internal/config"
	"go-test-inhousead/internal/handlers"
	"log"
	"net/http"
)

func main() {
	cfg := config.LoadConfig()

	app := handlers.NewApp(cfg)

	log.Fatal(http.ListenAndServe(cfg.ServerAddr, app))
}
