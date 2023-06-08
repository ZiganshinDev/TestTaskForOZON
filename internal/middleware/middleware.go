package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/ZiganshinDev/My-Pet-Projects/testForOzon/internal/database"
	"github.com/ZiganshinDev/My-Pet-Projects/testForOzon/internal/models"
	"github.com/gorilla/mux"
)

const (
	shortURLLength = 10
	chars          = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type response struct {
	ID       int64  `json:"id,omitempty"`
	Message  string `json:"message,omitempty"`
	ShortURL string `json:"url,omitempty"`
}

func GetShortURL(w http.ResponseWriter, r *http.Request) {
	var url models.URLs

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		log.Fatalf("Unable to decode the request body. %v", err)
	}

	if !isValidURL(url.OriginalURL) {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	insertID := database.InsertLongURL(url.OriginalURL)
	shortUrl := encode(insertID)

	res := response{
		ID:       insertID,
		Message:  "URL shortened successfully",
		ShortURL: fmt.Sprint(getProtocol(url.OriginalURL), shortUrl),
	}

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func GetLongURL(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	shortURL := params["shorturl"]

	id := decode(shortURL)

	decodedURL, err := database.GetLongURL(id)
	if err != nil {
		log.Fatalf("Unable to get url. %v", err)
	}

	w.Header().Set("Context-Type", "application/json")
	json.NewEncoder(w).Encode(decodedURL)
}

func encode(id int64) string {
	var code string
	for id > 0 {
		code = string(chars[id%62]) + code
		id /= 62
	}
	return code
}

func decode(code string) int64 {
	var id int64
	for _, c := range code {
		id = id*62 + int64(strings.IndexRune(string(chars), c))
	}
	return id
}

func isValidURL(input string) bool {
	u, err := url.Parse(input)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func getProtocol(url string) (protocol string) {
	protocol = strings.Split(url, "://")[0] + "://"

	return protocol
}
