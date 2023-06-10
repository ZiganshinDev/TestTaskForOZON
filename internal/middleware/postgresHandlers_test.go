package middleware_test

import (
	"testing"

	"github.com/ZiganshinDev/My-Pet-Projects/testForOzon/internal/middleware"
)

func TestGenerateShortURL(t *testing.T) {
	shortURL := middleware.GenerateShortURL()

	if len(shortURL) != middleware.ShortURLLength {
		t.Errorf("GenerateShortURL() = %q, want length %d", shortURL, middleware.ShortURLLength)
	}
}

func TestGetProtocol(t *testing.T) {
	protocol := middleware.GetProtocol("https://www.ozon.ru/")

	if protocol != "https://" {
		t.Errorf("GetProtocol() = %q, want https://", protocol)
	}
}
