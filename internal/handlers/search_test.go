package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"go-test-inhousead/internal/storage"
	"time"
)

func TestSearchHandler(t *testing.T) {
	storage.WebSitesCount = make(map[string]int)
	storage.WebSitesCount["/{id}"] = 0

	storage.WebSitesStore = make(map[string]time.Duration)
	storage.WebSitesStore["example.com"] = time.Second * 2
	storage.WebSitesStore["google.com"] = time.Second * 1

	r := chi.NewRouter()
	r.HandleFunc("/{id}", SearchHandler)

	req, err := http.NewRequest("GET", "/example.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидается статус %d, получен %d", http.StatusOK, status)
	}

	expectedResponse := "Сайт - example.com, время доступа к нему - 2s "
	if rr.Body.String() != expectedResponse {
		t.Errorf("Ожидается ответ %q, получен %q", expectedResponse, rr.Body.String())
	}

	if count := storage.WebSitesCount["/{id}"]; count != 1 {
		t.Errorf("Ожидается счетчик %d, получен %d", 1, count)
	}
}
