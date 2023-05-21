package service

import (
	"go-test-inhousead/internal/storage"
	"net/http"
	"time"
)

func PingHandler(_ http.ResponseWriter, _ *http.Request) {
	lists := storage.WebSitesStore

	for {
		Request(&lists)
		time.Sleep(time.Minute)
	}
}
