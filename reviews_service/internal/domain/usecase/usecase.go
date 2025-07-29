package usecase

import (
	"context"
	"reviews/internal/domain/entity"
	"reviews/internal/domain/repository"
)

type ReviewUsecase struct {
	Repo *repository.ReviewRepository
}

func NewReviewUsecase(repo *repository.ReviewRepository) *ReviewUsecase {
	return &ReviewUsecase{Repo: repo}
}

func (u *ReviewUsecase) AddReview(ctx context.Context, r *entity.Review) error {
	return u.Repo.Create(ctx, r)
}

func (u *ReviewUsecase) GetReviews(ctx context.Context, productID string) ([]entity.Review, error) {
	return u.Repo.GetByProductID(ctx, productID)
}

func (u *ReviewUsecase) DeleteReviews(ctx context.Context, productID string) error {
	return u.Repo.DeleteByProductID(ctx, productID)
}

func (u *ReviewUsecase) UpdateReview(ctx context.Context, r *entity.Review) error {
	return u.Repo.Update(ctx, r)
}
