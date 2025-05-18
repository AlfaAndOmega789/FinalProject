package main

import (
	"FinalProject/db"
	"FinalProject/routes"
	"log"
	"net/http"
)

func main() {
	db.InitDB()
	router := routes.SetupRouter()

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
