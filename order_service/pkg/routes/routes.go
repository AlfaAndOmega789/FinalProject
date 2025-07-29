package routes

import (
	"github.com/gorilla/mux"
	"order/internal/handler"
)

func SetupRouter(orderHandler *handler.OrderHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/orders/{id}", orderHandler.Get).Methods("GET")
	router.HandleFunc("/orders", orderHandler.Create).Methods("POST")
	router.HandleFunc("/orders/{id}", orderHandler.Delete).Methods("DELETE")
	router.HandleFunc("/orders/{id}", orderHandler.UpdateStatus).Methods("PATCH")

	return router
}
