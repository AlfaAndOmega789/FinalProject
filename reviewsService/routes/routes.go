package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"reviews/handlers"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/reviews", handlers.GetProducts).Methods("GET")

	return router
}
