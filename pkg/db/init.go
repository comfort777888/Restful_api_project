package db

import (
	"fmt"
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

var DB *sqlx.DB

func init() {
	DB = Newdb()
}

func Newdb() *sqlx.DB {
	connStr := "postgres://" + userName + ":" + password + "@" + host + ":" + port + "/" + dbname + "?" + "sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("error in connecting sql: %v", err)
	}
	// m, err := migrate.New("file://schema", connStr)
	// if err != nil {
	// 	log.Fatalf("failed to make migrate: %v", err)

	// }
	// if err := m.Up(); err != nil && err != migrate.ErrNoChange {
	// 	log.Fatal("failed to m.Up() in migrate")
	// }
	return db
}

func Migrate() {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", userName, password, host, port, dbname)

	m, err := migrate.New("file://migrate/", connStr)
	if err != nil {
		log.Printf("error migration: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("error migration up: %v", err)
	}
}
