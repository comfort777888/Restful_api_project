package users

import (
	"fmt"
	"log"
	"testing"

	"github.com/jmoiron/sqlx"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestUserCreate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection in TestUserCreate", err)
	}
	sqlxDB1 := sqlx.NewDb(db, "sqlmock")

	s1 := Store{DB: sqlxDB1}
	u1 := User{ID: 1, Data: "test_data_01"}
	mock.ExpectExec(`INSERT INTO users`).WithArgs(u1.ID, u1.Data).WillReturnResult(sqlmock.NewResult(1, 1))
	err = u1.Create(&s1)
	if err != nil {
		log.Printf("error in create test %v", err)
	}

	rows := mock.NewRows([]string{"id", "data"}).AddRow(1, "test_data_01")
	mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
	err = u1.Read(&s1)
	if err != nil {
		log.Fatalf("read %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		fmt.Println("got error:", err)
	}
}

func TestUserUpdate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection in TestUserUpdate", err)
	}
	sqlxDB1 := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB1.Close()
	s1 := Store{DB: sqlxDB1}
	u1 := User{ID: 1, Data: "test_data_01"}
	mock.ExpectExec(`INSERT INTO users`).WithArgs(u1.ID, u1.Data).WillReturnResult(sqlmock.NewResult(1, 1))
	err = u1.Create(&s1)
	if err != nil {
		log.Printf("error in create test %v", err)
	}

	rows := mock.NewRows([]string{"id", "data"}).AddRow(1, "test_data_01")
	mock.ExpectQuery(`SELECT`).WillReturnRows(rows)
	err = u1.Read(&s1)
	if err != nil {
		log.Fatalf("read %v", err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		fmt.Println("got error:", err)
	}
}
