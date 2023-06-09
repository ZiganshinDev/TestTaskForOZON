package middleware

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/ZiganshinDev/My-Pet-Projects/testForOzon/internal/database"
	"github.com/ZiganshinDev/My-Pet-Projects/testForOzon/internal/models"
	"github.com/gorilla/mux"
)

const (
	shortURLLength = 10
	chars          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
)

func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())

	var code strings.Builder
	for i := 0; i < shortURLLength; i++ {
		code.WriteByte(chars[rand.Intn(len(chars))])
	}
	return code.String()
}

type PostgreSQLService struct {
}

func (s *PostgreSQLService) GetShortURL(w http.ResponseWriter, r *http.Request) {
	var url models.URLs

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	if !isValidURL(url.OriginalURL) || database.NewPostgreSQLStorage().IsOriginalURLExists(url.OriginalURL) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	// Генерируем короткий URL и проверяем его на уникальность
	for {
		shortURL := generateShortURL()
		if !database.NewPostgreSQLStorage().IsShortURLExists(shortURL) {
			database.NewPostgreSQLStorage().InsertURLs(url.OriginalURL, shortURL)
			res := map[string]string{
				"short_url": getProtocol(url.OriginalURL) + shortURL,
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

	originalURL, err := database.NewPostgreSQLStorage().GetOriginalURL(shortURL)
	if err != nil {
		log.Fatalf("Unable to get url. %v", err)
	}

	res := map[string]string{
		"original_url": originalURL,
	}

	w.Header().Set("Context-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

func isValidURL(url string) bool {
	_, err := http.Get(url)
	return err == nil
}

func getProtocol(url string) (protocol string) {
	protocol = strings.Split(url, "://")[0] + "://"
	return protocol
}
