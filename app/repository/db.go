package repository

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func InitilizeDb(db *sql.DB, driverName string) (*sqlx.DB, error) {
	sqlxDb := sqlx.NewDb(db, driverName)
	err := sqlxDb.Ping()
	if err != nil {
		return nil, err
	}
	return sqlxDb, nil
}
