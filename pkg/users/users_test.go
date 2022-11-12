package users

import (
	"fmt"
	"log"
	"rest_api_project/pkg/db"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func TestUsersCRUD(t *testing.T) {
	db1, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection in TestUserCreate", err)
	}
	sqlxDB := sqlx.NewDb(db1, "sqlmock")
	db.DB = sqlxDB
	defer db.DB.Close()

	//Test for Create
	u1 := NewUser(1, "test_data_01")
	mock.ExpectExec(`INSERT INTO users`).WithArgs(u1.ID, u1.Data).WillReturnResult(sqlmock.NewResult(1, 1))
	err = u1.Create()
	if err != nil {
		log.Fatalf("error in create test %v", err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations in create: %s", err)
	}

	//Test for Read
	rows := mock.NewRows([]string{"id", "data"}).AddRow(1, "test_data_01")
	mock.ExpectQuery(`SELECT *`).WillReturnRows(rows)
	_, err = u1.All()
	if err != nil {
		log.Fatalf("read %v", err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		fmt.Println("got error:", err)
	}

	//Test for Update
	dataupd := "test_update"
	mock.ExpectExec(`UPDATE users`).WithArgs(dataupd, u1.ID).WillReturnResult(sqlmock.NewResult(0, 1))
	err = u1.Update(u1.ID, dataupd)
	if err != nil {
		log.Fatalf("error in update test %v", err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations in update: %s", err)
	}

	rows2 := mock.NewRows([]string{"id", "data"}).AddRow(1, "test_update")
	mock.ExpectQuery(`SELECT *`).WillReturnRows(rows2)
	_, err = u1.All()
	if err != nil {
		log.Fatalf("read %v", err)
	}
	if err = mock.ExpectationsWereMet(); err != nil {
		fmt.Println("got error:", err)
	}

	//Test for Delete

	mock.ExpectExec(`DELETE`).WithArgs(u1.ID).WillReturnResult(sqlmock.NewResult(0, 1))
	if err = u1.Delete(u1.ID); err != nil {
		log.Fatalf("error in delete test %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations in delete: %s", err)
	}
}
