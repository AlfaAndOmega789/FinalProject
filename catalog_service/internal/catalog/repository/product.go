package repository

import (
	"catalog/internal/catalog/entity"
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

func (r *ProductRepository) GetByID(id int) (entity.Product, error) {
	var product entity.Product
	err := r.DB.First(&product, id).Error
	return product, err
}

func (r *ProductRepository) Create(p entity.Product) (int, error) {
	if err := r.DB.Create(&p).Error; err != nil {
		return 0, err
	}
	return p.ID, nil
}

func (r *ProductRepository) Update(id int, p entity.Product) error {
	p.ID = id
	return r.DB.Save(&p).Error
}

func (r *ProductRepository) Delete(id int) error {
	return r.DB.Delete(&entity.Product{}, id).Error
}
