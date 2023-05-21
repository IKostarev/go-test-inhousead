package handlers

import (
	"fmt"
	"go-test-inhousead/internal/service"
	"go-test-inhousead/internal/storage"
	"net/http"
)

func LowHandler(w http.ResponseWriter, _ *http.Request) {
	count := storage.WebSitesCount
	count["/low"]++

	lists := storage.WebSitesStore

	url, low := service.SearchLowTime(&lists)

	answer := fmt.Sprintf(" Минимальное время доступа - %s к сайту - %s", low, url)

	_, _ = w.Write([]byte(answer))
}
