package migrations

import (
	"context"
	"log"
	"reviews_service/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func RunMigrations(db *mongo.Database) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	reviews := db.Collection("reviews")

	// Индекс
	indexModel := mongo.IndexModel{
		Keys: bson.M{"product_id": 1},
		Options: options.Index().
			SetBackground(true).
			SetName("idx_product_id"),
	}
	_, err := reviews.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Fatalf("Ошибка создания индекса: %v", err)
	}
	log.Println("Индекс по product_id создан")

	// Тестовые данные
	seed := []interface{}{
		models.Review{ProductID: "prod1", UserID: "user1", Rating: 5, Text: "Отлично!"},
		models.Review{ProductID: "prod1", UserID: "user2", Rating: 3, Text: "Нормально"},
	}
	_, err = reviews.InsertMany(ctx, seed)
	if err != nil {
		log.Fatalf("Ошибка наполнения: %v", err)
	}
	log.Println("Тестовые данные добавлены")
}
