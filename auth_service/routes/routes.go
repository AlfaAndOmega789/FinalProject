package routes

import (
	"auth/internal/user/handler"
	"github.com/gorilla/mux"
)

func SetupRouter(authHandler *handler.AuthHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
	r.HandleFunc("/auth/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/auth/me", authHandler.Me).Methods("GET")
	r.HandleFunc("/auth/refresh", authHandler.Refresh).Methods("POST")
	return r
}
