package middleware

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/ZiganshinDev/TestTaskForOZON/internal/database"
	"github.com/ZiganshinDev/TestTaskForOZON/internal/models"
	"github.com/gorilla/mux"
)

const (
	ShortURLLength = 10
	chars          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
)

func GenerateShortURL() string {
	rand.Seed(time.Now().UnixNano())

	var code strings.Builder
	for i := 0; i < ShortURLLength; i++ {
		code.WriteByte(chars[rand.Intn(len(chars))])
	}
	return code.String()
}

type PostgreSQLService struct {
	storage *database.PostgreSQLStorage
}

func NewPostgreSQLService(storage *database.PostgreSQLStorage) *PostgreSQLService {
	return &PostgreSQLService{
		storage: storage,
	}
}

func (s *PostgreSQLService) GetShortURL(w http.ResponseWriter, r *http.Request) {
	var url models.URLs

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	if !IsValidURL(url.OriginalURL) || s.storage.IsOriginalURLExists(url.OriginalURL) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	// Генерируем короткий URL и проверяем его на уникальность
	for {
		shortURL := GenerateShortURL()
		if !s.storage.IsShortURLExists(shortURL) {
			s.storage.InsertURLs(url.OriginalURL, shortURL)
			res := map[string]string{
				"short_url": GetProtocol(url.OriginalURL) + shortURL,
			}

			w.Header().Set("Context-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(res)

			break
		}
	}
}

func (s *PostgreSQLService) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	shortURL := params["shorturl"]

	originalURL, err := s.storage.GetOriginalURL(shortURL)
	if err != nil {
		log.Fatalf("Unable to get url. %v", err)
	}

	res := map[string]string{
		"original_url": originalURL,
	}

	w.Header().Set("Context-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func IsValidURL(url string) bool {
	_, err := http.Get(url)
	return err == nil
}

func GetProtocol(url string) (protocol string) {
	protocol = strings.Split(url, "://")[0] + "://"
	return protocol
}
