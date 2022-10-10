package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type product struct {
	Id           int
	ProductName  string
	Manufacturer string
	ProductCount string
	Price        int
}

type User struct {
	id   int
	data string
}

const (
	port     = "5432"
	userName = "postgres"
	password = "777888"
	host     = "localhost"
	dbname   = "postgres"
)

func main() {
	connStr := "postgres://" + userName + ":" + password + "@" + host + ":" + port + "/" + dbname + "?" + "sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error in opening sql: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("error in ping: %v", err)
	}

	m, err := migrate.New("file://schema", connStr)
	if err != nil {
		fmt.Println("failed to make migrate", err)
		return
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("failed to m.Up() in migrate")
		return
	}

	createTable(db)
	createUser(db)
	getProduct(db)
	getUser(db)
}

func createTable(db *sql.DB) {
	statement, err := db.Prepare("CREATE TABLE if NOT EXISTS products(Id SERIAL PRIMARY KEY, ProductName VARCHAR(30) NOT NULL, Manufacturer VARCHAR(20) NOT NULL, ProductCount INTEGER DEFAULT 0, Price NUMERIC);")
	if err != nil {
		log.Println(err)
	}
	statement.Exec()
}

func createUser(db *sql.DB) {
	statement, err := db.Prepare("INSERT INTO products VALUES (1, 'Galaxy S9', 'Samsung', 4, 63000);")
	if err != nil {
		log.Println(err)
	}
	statement.Exec()
}

func getProduct(db *sql.DB) {
	rows, err := db.Query("select * from products")
	if err != nil {
		log.Fatalf("error in select all rows: %v", err)
	}
	defer rows.Close()
	products := []product{}

	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.Id, &p.ProductName, &p.Manufacturer, &p.ProductCount, &p.Price)
		if err != nil {
			log.Printf("error in rows scan: %v", err)
			continue
		}
		products = append(products, p)
	}
	for _, p := range products {
		fmt.Println(p.Id, p.ProductName, p.Manufacturer, p.ProductCount, p.Price)
	}
}

func getUser(db *sql.DB) {
	rows, err := db.Query("select * from users")
	if err != nil {
		log.Fatalf("error in select all rows in Users: %v", err)
	}
	defer rows.Close()
	userss := []User{}

	for rows.Next() {
		p := User{}
		err := rows.Scan(&p.id, &p.data)
		if err != nil {
			log.Printf("error in rows scan User: %v", err)
			continue
		}
		userss = append(userss, p)
	}
	for _, p := range userss {
		fmt.Println(p.id, p.data)
	}
}
