package models

type URLs struct {
	ID          string `json:"urlid"`
	OriginalURL string `json:"original_url"`
	ShortURL    string `json:"short_url"`
}
