package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"

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
	if !IsValidURL(originalURL.URL) || valueExists(s.urlMap, originalURL.URL) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	// Генерируем случайный короткий URL
	shortURL := GenerateShortURL()
	for s.urlMap[shortURL] != "" {
		shortURL = GenerateShortURL()
	}

	// Сохраняем оригинальный и сокращенный URL в мапе
	s.urlMap[shortURL] = originalURL.URL

	// Отправляем клиенту сокращенный URL
	res := map[string]string{
		"short_url": GetProtocol(originalURL.URL) + shortURL,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

// Ищет оригинальный URL в мапе
func (s *InMemoryService) findOriginalURL(shortURL string) (string, error) {
	originalURL, ok := s.urlMap[shortURL]
	if !ok {
		return "", fmt.Errorf("URL not found")
	}
	return originalURL, nil
}

func (s *InMemoryService) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	shortURL := params["shorturl"]

	// Ищем оригинальный URL в мапе
	originalURL, err := s.findOriginalURL(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Отправляем клиенту оригинальный URL
	res := map[string]string{
		"original_url": originalURL,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func valueExists(m interface{}, value interface{}) bool {
	v := reflect.ValueOf(m)
	if v.Kind() != reflect.Map {
		return false
	}
	keys := v.MapKeys()
	for _, key := range keys {
		if reflect.DeepEqual(v.MapIndex(key).Interface(), value) {
			return true
		}
	}
	return false
}

func (s *InMemoryService) GetURLMap() map[string]string {
	return s.urlMap
}

func (s *InMemoryService) SetURLMap(urlMap map[string]string) {
	s.urlMap = urlMap
}
