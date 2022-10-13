package main

import (
	"cccccccc/pkg/users"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jmoiron/sqlx"
)

/*type product struct {
	Id           int
	ProductName  string
	Manufacturer string
	ProductCount string
	Price        int
}

type User struct {
	id   int
	data string
}*/

const (
	port     = "5432"
	userName = "postgres"
	password = "777888"
	host     = "localhost"
	dbname   = "postgres"
)

func main() {
	connStr := "postgres://" + userName + ":" + password + "@" + host + ":" + port + "/" + dbname + "?" + "sslmode=disable"

	/*db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("error in opening sql: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("error in ping: %v", err)
	}*/

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("error in connecting sql: %v", err)
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

	//CREATE
	u := users.NewUser(19, "Cat")
	err = u.Create(db)
	if err != nil {
		log.Fatalf("err crate user %v", err)
	}

	//READ
	fmt.Println("До удаления и изменения:")
	err = u.Read(db)
	if err != nil {
		log.Fatalf("err read user %v", err)
	}
	fmt.Println("last user:", u.Data)
	fmt.Println("После изменения и удаления:")
	err = u.Read(db)
	if err != nil {
		log.Fatalf("err read user №2 %v", err)
	}
	//UPDATE
	// обновляем строку с id=1
	_, err = db.Exec("UPDATE users SET data = 'Aray' where id = 1")
	if err != nil {
		panic(err)
	}
	//fmt.Println(result.RowsAffected()) // количество обновленных строк

	//DELETE
	// удаляем строку с id=2
	_, err = db.Exec("delete from users where id = $1", 2)
	if err != nil {
		panic(err)
	}
	//fmt.Println(result.RowsAffected()) // количество удаленных строк

	/*createTable(db)
		insertTable(db)
		getProduct(db)
		getUser(db)
	}

	func createTable(db *sqlx.DB) {
		statement, err := db.Preparex(`CREATE TABLE if NOT EXISTS products(Id SERIAL PRIMARY KEY, ProductName VARCHAR(30) NOT NULL, Manufacturer VARCHAR(20) NOT NULL, ProductCount INTEGER DEFAULT 0, Price NUMERIC);`)
		if err != nil {
			log.Printf("error in create table: %v", err)
		}
		statement.Exec()
	}

	func insertTable(db *sqlx.DB) {
		statement, err := db.Prepare("INSERT INTO products VALUES (1, 'Galaxy S9', 'Samsung', 4, 63000);")
		if err != nil {
			log.Printf("error in insert: %v", err)
		}
		statement.Exec()
	}

	func getProduct(db *sqlx.DB) {
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

	func getUser(db *sqlx.DB) {
		rows, err := db.Query("select * from userss")
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
		}*/
}
