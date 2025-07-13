package handlers

import (
	"net/http"
	"reviews_service/models"
	"reviews_service/usecase"

	"github.com/gin-gonic/gin"
)

type ReviewHandler struct {
	UC *usecase.ReviewUsecase
}

func NewReviewHandler(uc *usecase.ReviewUsecase) *ReviewHandler {
	return &ReviewHandler{UC: uc}
}

func (h *ReviewHandler) AddReview(c *gin.Context) {
	var r models.Review
	if err := c.ShouldBind(&r); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if err := h.UC.AddReview(c.Request.Context(), &r); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusCreated)
}

func (h *ReviewHandler) GetReviews(c *gin.Context) {
	productID := c.Param("product_id")
	reviews, err := h.UC.GetReviews(c.Request.Context(), productID)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, reviews)
}

func (h *ReviewHandler) DeleteReviews(c *gin.Context) {
	productID := c.Param("product_id")
	if err := h.UC.DeleteReviews(c.Request.Context(), productID); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}

func (h *ReviewHandler) UpdateReview(c *gin.Context) {
	var r models.Review
	if err := c.ShouldBind(&r); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if err := h.UC.UpdateReview(c.Request.Context(), &r); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Status(http.StatusOK)
}
