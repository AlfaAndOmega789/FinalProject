package entity

import (
	"time"

	"github.com/google/uuid"
)

type Currency struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Code      string    `gorm:"uniqueIndex;not null"`
	Rate      float64   `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}
