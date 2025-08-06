package routes

import (
	"reviews/internal/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoutes(r *gin.Engine, db *mongo.Database, h *handlers.ReviewHandler) {
	r.POST("/reviews", h.AddReview)
	r.GET("/reviews/:product_id", h.GetReviews)
	r.DELETE("/reviews/:product_id", h.DeleteReviews)
	r.PATCH("/reviews/:product_id", h.UpdateReview)
}
