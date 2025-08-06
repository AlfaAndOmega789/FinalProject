package usecase

import (
	"catalog/internal/domain/entity"
	"github.com/google/uuid"
)

type productUsecase struct {
	repo ProductRepository
}

func NewProductUsecase(repo ProductRepository) ProductUsecase {
	return &productUsecase{repo: repo}
}

func (uc *productUsecase) GetAll() ([]entity.Product, error) {
	return uc.repo.GetAll()
}

func (uc *productUsecase) GetByID(id uuid.UUID) (entity.Product, error) {
	return uc.repo.GetByID(id)
}

func (uc *productUsecase) Create(p entity.Product) (uuid.UUID, error) {
	return uc.repo.Create(p)
}

func (uc *productUsecase) Update(id uuid.UUID, p entity.Product) error {
	return uc.repo.Update(id, p)
}

func (uc *productUsecase) Delete(id uuid.UUID) error {
	return uc.repo.Delete(id)
}
