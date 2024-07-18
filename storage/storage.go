package storage

import (
	"database/sql"
)

func Setup() {
	sqldb, err := sql.Open(pgdriver)
}
