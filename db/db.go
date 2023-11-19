package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ConnectPostgres() *sql.DB {
	log.Println("Conexão com DB aberta")
	connStr := "user=postgres dbname=go-loja password=postgres host=localhost sslmode=disable port=5433"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panicln(err)
	}
	return db
}
