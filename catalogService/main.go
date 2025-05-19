package main

import (
	"catalog/db"
	"catalog/handlers"
	"catalog/routes"
	"log"
	"net/http"
)

func main() {
	database := db.InitDB()
	productHandler := &handlers.ProductHandler{DB: database}

	router := routes.SetupRouter(productHandler)

	log.Println("Сервер запущен на :8080")
	http.ListenAndServe(":8080", router)
}
