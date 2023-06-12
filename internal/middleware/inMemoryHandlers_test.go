package middleware_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ZiganshinDev/TestTaskForOZON/internal/middleware"
)

func TestGetShortURL(t *testing.T) {
	s := middleware.NewInMemoryService()

	// Проверка валидного запроса
	reqBody := map[string]string{
		"original_url": "https://www.ozon.ru/",
	}
	reqBodyBytes, _ := json.Marshal(reqBody)
	req, _ := http.NewRequest("POST", "/shorten", bytes.NewReader(reqBodyBytes))
	w := httptest.NewRecorder()
	s.GetShortURL(w, req)

	// Проверяем, что получили короткий URL
	var res map[string]string
	json.NewDecoder(w.Body).Decode(&res)
	if res["short_url"] == "" {
		t.Error("Expected short URL, but got empty response")
	}

	// Проверка невалидного запроса
	reqBody = map[string]string{
		"original_url": "invalid-url",
	}
	reqBodyBytes, _ = json.Marshal(reqBody)
	req, _ = http.NewRequest("POST", "/shorten", bytes.NewReader(reqBodyBytes))
	w = httptest.NewRecorder()
	s.GetShortURL(w, req)

	// Проверяем, что получили ошибку
	if w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code 400, but got %d", w.Code)
	}
}
