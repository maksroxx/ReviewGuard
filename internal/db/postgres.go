package db

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgreDB struct {
	Database *sql.DB
}

func NewPostgreDB(dsn string) *PostgreDB {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to PostgreSQL")
	return &PostgreDB{
		Database: db,
	}
}
