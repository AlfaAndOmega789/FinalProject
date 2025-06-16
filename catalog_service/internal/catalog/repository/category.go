package repository

import (
	"catalog/internal/catalog/entity"
	"gorm.io/gorm"
)

type CategoryPostgresRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *CategoryPostgresRepository {
	return &CategoryPostgresRepository{DB: db}
}

func (r *CategoryPostgresRepository) GetAll() ([]entity.Category, error) {
	var categories []entity.Category
	err := r.DB.Find(&categories).Error
	return categories, err
}
func (r *CategoryPostgresRepository) Create(p entity.Category) (int, error) {
	if err := r.DB.Create(&p).Error; err != nil {
		return 0, err
	}
	return p.ID, nil
}

func (r *CategoryPostgresRepository) Update(id int, p entity.Category) error {
	p.ID = id
	return r.DB.Save(&p).Error
}

func (r *CategoryPostgresRepository) Delete(id int) error {
	return r.DB.Delete(&entity.Category{}, id).Error
}
