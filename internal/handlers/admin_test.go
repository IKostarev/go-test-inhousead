package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdminHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/admin", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AdminHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Ожидается статус %d, получен %d", http.StatusOK, status)
	}
}
