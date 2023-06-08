package router

import (
	"github.com/ZiganshinDev/My-Pet-Projects/testForOzon/internal/middleware"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/shorten", middleware.GetShortURL).Methods("POST")
	router.HandleFunc("/shorten/db", middleware.GetShortURL).Methods("POST")
	router.HandleFunc("/original/{shorturl}", middleware.GetLongURL).Methods("GET")
	router.HandleFunc("/original/db/{shorturl}", middleware.GetLongURL).Methods("GET")

	return router
}
