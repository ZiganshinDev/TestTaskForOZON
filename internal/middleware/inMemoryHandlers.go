package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type InMemoryService struct {
	urlMap map[string]string // мапа для хранения оригинальных и сокращенных URL
}

func NewInMemoryService() *InMemoryService {
	return &InMemoryService{
		urlMap: make(map[string]string),
	}
}

func (s *InMemoryService) GetShortURL(w http.ResponseWriter, r *http.Request) {
	var originalURL struct {
		URL string `json:"original_url"`
	}

	err := json.NewDecoder(r.Body).Decode(&originalURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Проверяем, является ли URL валидным
	if !isValidURL(originalURL.URL) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	// Генерируем случайный короткий URL
	shortURL := generateShortURL()
	for s.urlMap[shortURL] != "" {
		shortURL = generateShortURL()
	}

	// Сохраняем оригинальный и сокращенный URL в мапе
	s.urlMap[shortURL] = originalURL.URL

	// Отправляем клиенту сокращенный URL
	res := map[string]string{
		"short_url": getProtocol(originalURL.URL) + shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func (s *InMemoryService) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	shortURL := params["shorturl"]

	// Ищем оригинальный URL в мапе
	originalURL, ok := s.urlMap[shortURL]
	if !ok {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Отправляем клиенту оригинальный URL
	res := map[string]string{
		"original_url": originalURL,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
