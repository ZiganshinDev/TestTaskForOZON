package app

import (
	"log"
	"net/http"
	"os"

	"github.com/ZiganshinDev/My-Pet-Projects/testForOzon/internal/router"
)

func Run() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatalf("Error loading .env file. %v", err)
	// }

	storageType := os.Getenv("STORAGE_TYPE")

	log.Println("Using storage:", storageType)

	if storageType == "in-memory" {
		// Инициализация in-memory хранилища
		r := router.Router("in-memory")

		log.Println("Starting server on the port 8080...")

		log.Fatal(http.ListenAndServe(":8080", r))
	} else if storageType == "postgres" {
		// Инициализация PostgreSQL хранилища
		r := router.Router("postgres")

		log.Println("Starting server on the port 8080...")

		log.Fatal(http.ListenAndServe(":8080", r))
	} else {
		// Обработка некорректного значения флага
		log.Fatal("Unknown storage type:", storageType)
		return
	}
}
