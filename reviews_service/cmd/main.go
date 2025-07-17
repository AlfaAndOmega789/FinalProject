package main

import (
	"log"
	"net/http"
	"os"
	"reviews/db"
	"reviews/handlers"
	"reviews/migrations"
	"reviews/repository"
	"reviews/routes"
	"reviews/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitMongo()
	database := db.Client.Database("reviews_db")

	migrations.RunMigrations(database)

	repo := repository.NewReviewRepository(database)
	uc := usecase.NewReviewUsecase(repo)
	handler := handlers.NewReviewHandler(uc)

	router := gin.Default()
	routes.InitRoutes(router, database, handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	log.Println("Reviews Service запущен на порту :" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
