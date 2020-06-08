package models

import "database/sql"

type AppDb struct {
	DB *sql.DB
}
