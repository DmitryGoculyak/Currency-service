package db

import (
	"github.com/jmoiron/sqlx"
	"log"

	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func InitDB() {
	var err error
	conn := "postgres://postgres:2020@localhost/currancy?sslmode=disable"
	DB, err = sqlx.Connect("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Println("[DATABASE] Successfully connected!")
}
