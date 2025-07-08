package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:heru12345@localhost:5432/client?sslmode=disable")
	if err != nil {
		log.Fatalf("error init db: %v\n", err)
	}

	// Tes koneksi
	if err = db.Ping(); err != nil {
		log.Fatalf("tidak bisa konek ke database: %v\n", err)
	}

	return db
}
