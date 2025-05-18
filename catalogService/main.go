package main

import (
	"catalog/routes"
	"log"
	"net/http"
)

func main() {

	router := routes.SetupRouter()

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
