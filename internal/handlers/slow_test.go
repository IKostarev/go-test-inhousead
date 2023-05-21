package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go-test-inhousead/internal/storage"
	"time"
)

func TestSlowHandler(t *testing.T) {
	storage.WebSitesCount = make(map[string]int)
	storage.WebSitesCount["/slow"] = 0

	storage.WebSitesStore = make(map[string]time.Duration)
	storage.WebSitesStore["example.com"] = time.Second * 2
	storage.WebSitesStore["google.com"] = time.Second * 3

	r := http.NewServeMux()
	r.HandleFunc("/slow", SlowHandler)

	req, err := http.NewRequest("GET", "/slow", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидается статус %d, получен %d", http.StatusOK, status)
	}

	expectedResponse := " Максимальное время доступа - 3s к сайту - google.com"
	if rr.Body.String() != expectedResponse {
		t.Errorf("Ожидается ответ %q, получен %q", expectedResponse, rr.Body.String())
	}

	if count := storage.WebSitesCount["/slow"]; count != 1 {
		t.Errorf("Ожидается счетчик %d, получен %d", 1, count)
	}
}
