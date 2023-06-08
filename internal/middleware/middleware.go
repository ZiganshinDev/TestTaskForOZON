package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
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

type response struct {
	ID       int64  `json:"id,omitempty"`
	Message  string `json:"message,omitempty"`
	ShortURL string `json:"url,omitempty"`
}

func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())

	var code strings.Builder
	for i := 0; i < shortURLLength; i++ {
		code.WriteByte(chars[rand.Intn(len(chars))])
	}
	return code.String()
}

func GetShortURL(w http.ResponseWriter, r *http.Request) {
	var url models.URLs

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	if !isValidURL(url.OriginalURL) || !database.IsOriginalURLExists(url.OriginalURL) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	// Генерируем короткий URL и проверяем его на уникальность
	for {
		shortURL := generateShortURL()
		if !database.IsShortURLExists(shortURL) {
			insertID := database.InsertURLs(url.OriginalURL, shortURL)
			res := response{
				ID:       insertID,
				Message:  "URL shortened successfully",
				ShortURL: fmt.Sprint(getProtocol(url.OriginalURL), shortURL),
			}

			w.Header().Set("Context-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(res)

			break
		}
	}
}

func GetLongURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	shortURL := params["shorturl"]

	originalURL, err := database.GetLongURL(shortURL)
	if err != nil {
		log.Fatalf("Unable to get url. %v", err)
	}

	w.Header().Set("Context-Type", "application/json")
	json.NewEncoder(w).Encode(originalURL)
}

func isValidURL(input string) bool {
	u, err := url.Parse(input)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func getProtocol(url string) (protocol string) {
	protocol = strings.Split(url, "://")[0] + "://"
	return protocol
}
