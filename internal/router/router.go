package router

import (
	"github.com/ZiganshinDev/My-Pet-Projects/testForOzon/internal/database"
	"github.com/ZiganshinDev/My-Pet-Projects/testForOzon/internal/middleware"
	"github.com/gorilla/mux"
)

func Router(storage string) *mux.Router {
	router := mux.NewRouter()

	if storage == "postgres" {
		db := database.NewPostgreSQLStorage()
		service := middleware.NewPostgreSQLService(db)

		router.HandleFunc("/short", service.GetShortURL).Methods("POST")
		router.HandleFunc("/original/{shorturl}", service.GetOriginalURL).Methods("GET")
	} else {
		urlService := middleware.NewInMemoryService()

		router.HandleFunc("/short", urlService.GetShortURL).Methods("POST")
		router.HandleFunc("/original/{shorturl}", urlService.GetOriginalURL).Methods("GET")
	}

	return router
}
