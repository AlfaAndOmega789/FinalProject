package models

import "database/sql"

type Roles struct {
	ID           sql.NullString
	Name         string //должно быть уникальным
	Descriptions string
}
