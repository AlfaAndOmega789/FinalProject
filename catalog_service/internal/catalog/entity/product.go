package entity

import (
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	CategoryID  int
	CreatedAt   time.Time
}
