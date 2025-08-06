package usecase

import (
	"catalog/internal/domain/entity"
	"github.com/google/uuid"
)

type categoryUsecase struct {
	repo CategoryRepository
}

func NewCategoryUsecase(repo CategoryRepository) CategoryUsecase {
	return &categoryUsecase{repo: repo}
}

func (uc *categoryUsecase) GetAll() ([]entity.Category, error) {
	return uc.repo.GetAll()
}

func (uc *categoryUsecase) Create(p entity.Category) (uuid.UUID, error) {
	return uc.repo.Create(p)
}

func (uc *categoryUsecase) Update(id uuid.UUID, p entity.Category) error {
	return uc.repo.Update(id, p)
}

func (uc *categoryUsecase) Delete(id uuid.UUID) error {
	return uc.repo.Delete(id)
}
