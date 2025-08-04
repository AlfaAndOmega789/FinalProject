package postgres

import (
	"background/internal/domain/entity"
	"background/internal/domain/repository"
	"gorm.io/gorm"
)

type CurrencyGormRepository struct {
	db *gorm.DB
}

func NewCurrencyGormRepository(db *gorm.DB) repository.CurrencyRepository {
	return &CurrencyGormRepository{db: db}
}

func (r *CurrencyGormRepository) SaveBatch(currencies []entity.Currency) error {
	return r.db.Save(&currencies).Error
}
