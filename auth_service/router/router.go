package router

import (
	"github.com/gorilla/mux"
)

func SetupRouter(registerHandler *handler.RegisterHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/auth/register", registerHandler.Register).Methods("POST")
	return r
}
