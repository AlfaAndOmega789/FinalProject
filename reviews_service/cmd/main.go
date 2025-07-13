package main

import (
	"log"
	"net/http"
	"os"
	"reviews_service/db"
	"reviews_service/handlers"
	"reviews_service/migrations"
	"reviews_service/repository"
	"reviews_service/routes"
	"reviews_service/usecase"

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
		port = "8082"
	}

	log.Println("Reviews Service запущен на порту :" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
