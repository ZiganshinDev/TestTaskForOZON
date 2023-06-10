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
	// Создаем роутер
	r := mux.NewRouter()

	// Создаем объект InMemoryService и регистрируем на роутере обработчики
	s := middleware.NewInMemoryService()
	r.HandleFunc("/shorten", s.GetShortURL).Methods("POST")
	r.HandleFunc("/{shorturl}", s.GetOriginalURL).Methods("GET")

	// Создаем тестовый запрос
	reqBody := map[string]string{
		"original_url": "https://www.ozon.ru/",
	}
	reqBodyJSON, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/shorten", bytes.NewBuffer(reqBodyJSON))
	w := httptest.NewRecorder()

	// Вызываем метод GetShortURL
	r.ServeHTTP(w, req)

	// Проверяем, что ответ содержит корректный сокращенный URL
	var resBody map[string]string
	json.NewDecoder(w.Body).Decode(&resBody)
	shortURL := resBody["short_url"]
	if !middleware.IsValidURL(shortURL) || !strings.HasPrefix(shortURL, "http") {
		t.Errorf("GetShortURL() = %q, want valid URL", shortURL)
	}

	// Проверяем, что мапа urlMap содержит запись для созданного короткого URL
	urlMap := s.GetURLMap()
	if len(urlMap) != 1 || urlMap[shortURL] != "https://www.ozon.ru/" {
		t.Errorf("urlMap = %v, want %v", urlMap, map[string]string{shortURL: "https://www.ozon.ru/"})
	}
}

func TestGetOriginalURL(t *testing.T) {
	// Создаем роутер
	r := mux.NewRouter()

	// Создаем объект InMemoryService и регистрируем на роутере обработчики
	s := middleware.NewInMemoryService()
	r.HandleFunc("/shorten", s.GetShortURL).Methods("POST")
	r.HandleFunc("/{shorturl}", s.GetOriginalURL).Methods("GET")

	// Создаем тестовый запрос
	req := httptest.NewRequest("GET", "/abc123", nil)
	req = mux.SetURLVars(req, map[string]string{"shorturl": "abc123"})
	w := httptest.NewRecorder()

	// Устанавливаем мапу urlMap для объекта InMemoryService
	s.SetURLMap(map[string]string{"abc123": "https://www.ozon.ru/"})

	// Вызываем метод GetOriginalURL
	r.ServeHTTP(w, req)

	// Проверяем, что ответ содержит корректный оригинальный URL
	var resBody map[string]string
	json.NewDecoder(w.Body).Decode(&resBody)
	originalURL := resBody["original_url"]
	if originalURL != "https://www.ozon.ru/" {
		t.Errorf("GetOriginalURL() = %q, want %q", originalURL, "https://www.ozon.ru/")
	}
}
