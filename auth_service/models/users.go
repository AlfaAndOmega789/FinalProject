package models

import (
	"database/sql"
	"time"
)

type Users struct {
	ID           string
	Email        string
	PasswordHash string
	Name         string
	RoleID       sql.NullInt64
	CreatedAt    time.Time
}
