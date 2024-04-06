package database

import (
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func ConnectDB(DBURL string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("pgx", DBURL)
	db.SetMaxOpenConns(1000) // The default is 0 (unlimited)
	db.SetMaxIdleConns(10)   // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(0)

	if err != nil {
		return nil, err
	}
	return db, nil
}
