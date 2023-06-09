package router

import (
	"github.com/ZiganshinDev/My-Pet-Projects/testForOzon/internal/middleware"
	"github.com/gorilla/mux"
)

func Router(storage string) *mux.Router {
	router := mux.NewRouter()

	if storage == "postgres" {
		data := middleware.PostgreSQLService{}

		router.HandleFunc("/shorten", data.GetShortURL).Methods("POST")
		router.HandleFunc("/original/{shorturl}", data.GetOriginalURL).Methods("GET")
	} else {
		urlService := middleware.NewInMemoryService()

		router.HandleFunc("/shorten", urlService.GetShortURL).Methods("POST")
		router.HandleFunc("/original/{shorturl}", urlService.GetOriginalURL).Methods("GET")
	}

	return router
}
