package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"go-test-inhousead/internal/storage"
	"net/http"
	"time"
)

func search(s string, list *map[string]time.Duration) (time.Duration, bool) {
	for url, dur := range *list {
		if url == s && dur > 0 {
			return dur, true
		}
	}

	return 0, false
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	count := storage.WebSitesCount
	count["/{id}"]++

	url := chi.URLParam(r, "id")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		_ = fmt.Errorf("URLParam is empty")
		return
	}

	store := storage.WebSitesStore

	dur, ok := search(url, &store)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		_ = fmt.Errorf("ВВЕДЕННЫЙ ВАМИ url - %s НЕ НАЙДЕН ЛИБО К НЕМУ НЕТ ДОСТУПА", url)
		return
	}

	answer := fmt.Sprintf("Сайт - %s, время доступа к нему - %s ", url, dur)

	_, _ = w.Write([]byte(answer))
}
