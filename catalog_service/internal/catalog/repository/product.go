package repository

import (
	"catalog/internal/catalog/entity"
	"gorm.io/gorm"
)

type ProductPostgresRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductPostgresRepository {
	return &ProductPostgresRepository{DB: db}
}

func (r *ProductPostgresRepository) GetAll() ([]entity.Product, error) {
	var products []entity.Product
	err := r.DB.Find(&products).Error
	return products, err
}

func (r *ProductPostgresRepository) GetByID(id int) (entity.Product, error) {
	var product entity.Product
	err := r.DB.First(&product, id).Error
	return product, err
}

func (r *ProductPostgresRepository) Create(p entity.Product) (int, error) {
	if err := r.DB.Create(&p).Error; err != nil {
		return 0, err
	}
	return p.ID, nil
}

func (r *ProductPostgresRepository) Update(id int, p entity.Product) error {
	p.ID = id
	return r.DB.Save(&p).Error
}

func (r *ProductPostgresRepository) Delete(id int) error {
	return r.DB.Delete(&entity.Product{}, id).Error
}
