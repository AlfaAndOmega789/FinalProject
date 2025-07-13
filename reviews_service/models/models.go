package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Review struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	ProductID string             `bson:"product_id"`
	UserID    string             `bson:"user_id"`
	Rating    int                `bson:"rating"`
	Text      string             `bson:"text"`
}
