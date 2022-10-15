package db

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	port     = "5432"
	userName = "postgres"
	password = "777888"
	host     = "localhost"
	dbname   = "postgres"
)

func init() {
	Newdb()
}

func Newdb() *sqlx.DB {
	connStr := "postgres://" + userName + ":" + password + "@" + host + ":" + port + "/" + dbname + "?" + "sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("error in connecting sql: %v", err)
	}
	m, err := migrate.New("file://schema", connStr)
	if err != nil {
		log.Fatalf("failed to make migrate: %v", err)

	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to m.Up() in migrate")
	}
	return db
}
