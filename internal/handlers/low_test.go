package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go-test-inhousead/internal/storage"
)

func TestLowHandler(t *testing.T) {
	storage.WebSitesCount = make(map[string]int)
	storage.WebSitesCount["/low"] = 0

	storage.WebSitesStore = make(map[string]time.Duration)
	storage.WebSitesStore["example.com"] = time.Second * 2
	storage.WebSitesStore["google.com"] = time.Second * 1

	req, err := http.NewRequest("GET", "/low", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(LowHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидается статус %d, получен %d", http.StatusOK, status)
	}

	expectedResponse := " Минимальное время доступа - 1s к сайту - google.com"
	if rr.Body.String() != expectedResponse {
		t.Errorf("Ожидается ответ %q, получен %q", expectedResponse, rr.Body.String())
	}

	if count := storage.WebSitesCount["/low"]; count != 1 {
		t.Errorf("Ожидается счетчик %d, получен %d", 1, count)
	}
}
