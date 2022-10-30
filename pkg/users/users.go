package users

import (
	"fmt"
	"log"
	"rest_api_project/pkg/db"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID   int    `json:"id" db:"id"`
	Data string `json:"data" db:"data"`
}

type Store struct {
	DB *sqlx.DB
}

func NewStore(db *sqlx.DB) *Store {
	return &Store{
		DB: db,
	}
}

func NewUser(id int, data string) *User {
	return &User{
		ID:   id,
		Data: data,
	}
}

func (u *User) Create() error {
	sqlStr := "INSERT INTO users (id, data) VALUES (:id, :data)"
	_, err := db.DB.NamedExec(sqlStr,
		map[string]interface{}{
			"id":   u.ID,
			"data": u.Data,
		})
	return err
}

func (u *User) Read() error {
	rows, err := db.DB.Queryx("SELECT * FROM users")
	for rows.Next() {
		err := rows.StructScan(&u)
		if err != nil {
			fmt.Printf("scan failed, error:%v\n", err)
			continue
		}
		fmt.Println(u.ID, u.Data)
	}
	return err
}

func (u *User) Update(id int, data string) error {
	sqlStr := "UPDATE users SET data=:data WHERE id=:id"
	_, err := db.DB.NamedExec(sqlStr,
		map[string]interface{}{
			"id":   id,
			"data": data,
		})
	return err
}

func (u *User) Delete(id int) error {
	sqlStr := "DELETE FROM users WHERE id =:id"
	_, err := db.DB.NamedExec(sqlStr,
		map[string]interface{}{
			"id": id,
		})
	return err
}

func (u *User) ReadRow(id int) error {
	rows, err := db.DB.NamedQuery(`SELECT * FROM users WHERE id=:fn`, map[string]interface{}{"fn": id})
	for rows.Next() {
		err := rows.StructScan(&u)
		if err != nil {
			log.Fatalf("error in readrow scanning: %v", err)
		}
		fmt.Printf("Row with selected ID: %+v", u)
	}
	return err
}
