package entity

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"not null"`
	CategoryID  uuid.UUID `gorm:"type:uuid"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
