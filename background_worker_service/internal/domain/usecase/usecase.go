package usecase

import (
	"background/internal/domain/entity"
	"background/internal/domain/repository"
)

type CurrencyUseCase struct {
	repo repository.CurrencyRepository
}

func NewCurrencyUseCase(repo repository.CurrencyRepository) *CurrencyUseCase {
	return &CurrencyUseCase{repo: repo}
}

func (uc *CurrencyUseCase) SaveRates(rates []entity.Currency) error {
	return uc.repo.SaveBatch(rates)
}
