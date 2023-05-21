package handlers

import (
	"fmt"
	"go-test-inhousead/internal/storage"
	"net/http"
)

func AdminHandler(w http.ResponseWriter, _ *http.Request) {
	count := storage.WebSitesCount

	for i, j := range count {
		answer := fmt.Sprintf("«Панель администратора» ендпоинт - %s, количество обращений - %d \n", i, j)
		_, _ = w.Write([]byte(answer))
	}
}
