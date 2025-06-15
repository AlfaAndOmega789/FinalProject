package models

import "database/sql"

type Roles struct {
	ID           sql.NullString
	Name         string
	Descriptions string
}
