package entity

import (
	"github.com/google/uuid"
	"time"
)

type Category struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
