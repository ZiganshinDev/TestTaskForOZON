package middleware_test

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ZiganshinDev/My-Pet-Projects/testForOzon/internal/middleware"
	"github.com/gorilla/mux"
)

func TestGetShortURL(t *testing.T) {
	// Создаем тестовый запрос
	reqBody := map[string]string{
		"original_url": "https://example.com",
	}
	reqBodyJSON, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/shorten", bytes.NewBuffer(reqBodyJSON))
	w := httptest.NewRecorder()

	// Создаем объект InMemoryService
	s := middleware.NewInMemoryService()

	// Вызываем метод GetShortURL
	s.GetShortURL(w, req)

	// Проверяем, что ответ содержит корректный сокращенный URL
	var resBody map[string]string
	json.NewDecoder(w.Body).Decode(&resBody)
	shortURL := resBody["short_url"]
	if !middleware.IsValidURL(shortURL) || !strings.HasPrefix(shortURL, "http") {
		t.Errorf("GetShortURL() = %q, want valid URL", shortURL)
	}
}

func TestGetOriginalURL(t *testing.T) {
	// Создаем тестовый запрос
	req := httptest.NewRequest("GET", "/abc123", nil)
	req = mux.SetURLVars(req, map[string]string{"shorturl": "abc123"})
	w := httptest.NewRecorder()

	// Создаем объект InMemoryService
	s := middleware.NewInMemoryService()

	// Вызываем метод GetOriginalURL
	s.urlMap["abc123"] = "https://example.com"
	s.GetOriginalURL(w, req)

	// Проверяем, что ответ содержит корректный оригинальный URL
	var resBody map[string]string
	json.NewDecoder(w.Body).Decode(&resBody)
	originalURL := resBody["original_url"]
	if originalURL != "https://example.com" {
		t.Errorf("GetOriginalURL() = %q, want %q", originalURL, "https://example.com")
	}
}
