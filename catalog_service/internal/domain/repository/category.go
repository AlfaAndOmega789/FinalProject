package repository

import (
	"catalog/internal/domain/entity"
	"github.com/google/uuid"
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

func (r *CategoryPostgresRepository) Create(p entity.Category) (uuid.UUID, error) {
	if err := r.DB.Create(&p).Error; err != nil {
		return uuid.Nil, err
	}
	return p.ID, nil
}

func (r *CategoryPostgresRepository) Update(id uuid.UUID, p entity.Category) error {
	p.ID = id
	return r.DB.Save(&p).Error
}

func (r *CategoryPostgresRepository) Delete(id uuid.UUID) error {
	return r.DB.Delete(&entity.Category{}, "id = ?", id).Error
}
