package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"reviews/internal/domain/entity"
)

type ReviewRepository struct {
	Collection *mongo.Collection
}

func NewReviewRepository(db *mongo.Database) *ReviewRepository {
	return &ReviewRepository{
		Collection: db.Collection("reviews"),
	}
}

func (r *ReviewRepository) Create(ctx context.Context, review *entity.Review) error {
	_, err := r.Collection.InsertOne(ctx, review)
	return err
}

func (r *ReviewRepository) GetByProductID(ctx context.Context, productID string) ([]entity.Review, error) {
	cursor, err := r.Collection.Find(ctx, bson.M{"product_id": productID})
	if err != nil {
		return nil, err
	}

	var reviews []entity.Review
	if err := cursor.All(ctx, &reviews); err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *ReviewRepository) DeleteByProductID(ctx context.Context, productID string) error {
	_, err := r.Collection.DeleteMany(ctx, bson.M{"product_id": productID})
	return err
}

func (r *ReviewRepository) Update(ctx context.Context, review *entity.Review) error {
	filter := bson.M{"product_id": review.ProductID, "user_id": review.UserID}
	update := bson.M{"$set": review}
	_, err := r.Collection.UpdateOne(ctx, filter, update)
	return err
}
