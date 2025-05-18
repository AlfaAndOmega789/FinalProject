package models

import (
	"database/sql"
	"time"
)

type Product struct {
	ID          int
	Name        string
	Description sql.NullString
	Price       float64
	CategoryID  sql.NullInt64 // если category_id тоже может быть NULL
	CreatedAt   time.Time
}
