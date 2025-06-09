package usecase

import (
	"catalog/internal/catalog/entity"
)

type ProductRepository interface {
	GetAll() ([]entity.Product, error)
	GetByID(id int) (entity.Product, error)
	Create(p entity.Product) (int, error)
	Update(id int, p entity.Product) error
	Delete(id int) error
}

type ProductUsecase interface {
	GetAll() ([]entity.Product, error)
	GetByID(id int) (entity.Product, error)
	Create(p entity.Product) (int, error)
	Update(id int, p entity.Product) error
	Delete(id int) error
}
