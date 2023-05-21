package handlers

import (
	"fmt"
	"go-test-inhousead/internal/service"
	"go-test-inhousead/internal/storage"
	"net/http"
)

func SlowHandler(w http.ResponseWriter, _ *http.Request) {
	count := storage.WebSitesCount
	count["/slow"]++

	lists := storage.WebSitesStore

	url, slow := service.SearchSlowTime(&lists)

	answer := fmt.Sprintf(" Максимальное время доступа - %s к сайту - %s", slow, url)

	_, _ = w.Write([]byte(answer))
}
