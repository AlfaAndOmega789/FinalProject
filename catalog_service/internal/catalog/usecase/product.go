package usecase

import (
	"catalog/internal/catalog/entity"
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

func (uc *productUsecase) GetByID(id int) (entity.Product, error) {
	return uc.repo.GetByID(id)
}

func (uc *productUsecase) Create(p entity.Product) (int, error) {
	return uc.repo.Create(p)
}

func (uc *productUsecase) Update(id int, p entity.Product) error {
	return uc.repo.Update(id, p)
}

func (uc *productUsecase) Delete(id int) error {
	return uc.repo.Delete(id)
}
