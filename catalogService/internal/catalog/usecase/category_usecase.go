package usecase

import (
	"catalog/internal/catalog/entity"
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

func (uc *categoryUsecase) Create(p entity.Category) (int, error) {
	return uc.repo.Create(p)
}

func (uc *categoryUsecase) Update(id int, p entity.Category) error {
	return uc.repo.Update(id, p)
}

func (uc *categoryUsecase) Delete(id int) error {
	return uc.repo.Delete(id)
}
