package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"reviews/internal/domain/repository"
	"reviews/internal/domain/usecase"
	"reviews/internal/handlers"
	"reviews/internal/infrastructure/mongo"
	"reviews/internal/infrastructure/routes"
	"reviews/migrations"
)

func main() {
	mongo.InitMongo()
	database := mongo.Client.Database("reviews_db")

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
