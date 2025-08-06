package repository

import "background/internal/domain/entity"

type CurrencyRepository interface {
	SaveBatch([]entity.Currency) error
}
