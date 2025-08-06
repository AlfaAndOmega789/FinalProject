package usecase

import (
	"catalog/internal/domain/entity"
	"github.com/google/uuid"
)

type ProductRepository interface {
	GetAll() ([]entity.Product, error)
	GetByID(id uuid.UUID) (entity.Product, error)
	Create(p entity.Product) (uuid.UUID, error)
	Update(id uuid.UUID, p entity.Product) error
	Delete(id uuid.UUID) error
}

type ProductUsecase interface {
	GetAll() ([]entity.Product, error)
	GetByID(id uuid.UUID) (entity.Product, error)
	Create(p entity.Product) (uuid.UUID, error)
	Update(id uuid.UUID, p entity.Product) error
	Delete(id uuid.UUID) error
}

type CategoryRepository interface {
	GetAll() ([]entity.Category, error)
	Create(p entity.Category) (uuid.UUID, error)
	Update(id uuid.UUID, p entity.Category) error
	Delete(id uuid.UUID) error
}

type CategoryUsecase interface {
	GetAll() ([]entity.Category, error)
	Create(p entity.Category) (uuid.UUID, error)
	Update(id uuid.UUID, p entity.Category) error
	Delete(id uuid.UUID) error
}
