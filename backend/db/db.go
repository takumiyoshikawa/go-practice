package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	db *sql.DB
}

func NewDatabase() (*DB, error) {
	db, err := sql.Open("postgres", "postgres://root:password@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}
	
	return &DB{db: db}, nil
}

func (d *DB) Close() {
	d.db.Close()
}

func (d *DB) GetDB() *sql.DB {
	return d.db
}