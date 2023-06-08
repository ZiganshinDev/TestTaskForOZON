package app

import (
	"log"
	"net/http"

	"github.com/ZiganshinDev/My-Pet-Projects/testForOzon/internal/router"
	"github.com/joho/godotenv"
)

func Run() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file. %v", err)
	}

	r := router.Router()

	log.Println("Starting server on the port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
