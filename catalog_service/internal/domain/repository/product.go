package repository

import (
	"catalog/internal/domain/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	err := r.DB.Find(&products).Error
	return products, err
}

func (r *ProductRepository) GetByID(id uuid.UUID) (entity.Product, error) {
	var product entity.Product
	err := r.DB.First(&product, "id = ?", id).Error
	return product, err
}

func (r *ProductRepository) Create(p entity.Product) (uuid.UUID, error) {
	if err := r.DB.Create(&p).Error; err != nil {
		return uuid.Nil, err
	}
	return p.ID, nil
}

func (r *ProductRepository) Update(id uuid.UUID, p entity.Product) error {
	p.ID = id
	return r.DB.Save(&p).Error
}

func (r *ProductRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&entity.Product{}, "id = ?", id).Error
}
