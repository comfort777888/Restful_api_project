package db

import (
	"fmt"
	"log"

	"github.com/pressly/goose/v3"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	//_ "github.com/golang-migrate/migrate/v4/database/postgres"
	//_ "github.com/golang-migrate/migrate/v4/source/file"
)

var DB *sqlx.DB

//var embedMigrations embed.FS

func init() {
	DB = Newdb()
	//Migrate(DB)
}

func Configs() string {
	viper.SetConfigName("development")
	viper.AddConfigPath("./")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("error viper get port: %v", err)
	}
	slice := fmt.Sprintf("user=%v dbname=%s host=%s port=%s password=%s sslmode=%s",
		viper.Get("services.postgres.userName"),
		viper.Get("services.postgres.dbname"),
		viper.Get("services.postgres.host"),
		viper.Get("services.postgres.port"),
		viper.Get("services.postgres.password"),
		viper.Get("services.postgres.sslmode"))
	fmt.Println(slice)
	return slice
}

func Newdb() *sqlx.DB {
	//connStr := "postgres://" + userName + ":" + password + "@" + host + ":" + port + "/" + dbname + "?" + "sslmode=disable"
	connStr := Configs()
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("error in connecting sql: %v", err)
	}

	//goose.SetBaseFS(embedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("error in goose SetDialect:%v", err)
	}

	if err := goose.Up(db.DB, "."); err != nil {
		log.Fatalf("error in goose Up:%v", err)
	}
	// db.Exec(`CREATE TABLE IF NOT EXISTS users (
	// 	id int primary key,
	// 	data VARCHAR
	// );`)

	return db
}

//func Migrate(DB *sqlx.DB) {
// 	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", userName, password, host, port, dbname)

// 	m, err := migrate.New("file://schema", connStr)
// 	if err != nil {
// 		log.Fatalf("failed to make migrate: %v", err)
// 	}

// 	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
// 		log.Fatalf("failed to m.Up() in migrate: %v", err)
// 	}

//}
