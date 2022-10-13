package users

import (
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID   int
	Data string
}

func NewUser(id int, data string) *User {
	return &User{
		ID:   id,
		Data: data,
	}
}

func (u *User) Create(db *sqlx.DB) error {
	_, err := db.NamedExec(`INSERT INTO users (id, data) VALUES (:id, :data)`,
		map[string]interface{}{
			"id":   u.ID,
			"data": u.Data,
		})
	return err
}

func (u *User) Read(db *sqlx.DB) error {
	place := User{}
	rows, err := db.Queryx("SELECT * FROM users")
	if err != nil {
		log.Printf("err in sel:%v", err)
	}

	for rows.Next() {
		err := rows.StructScan(&place)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", place)
	}

	row, err := db.NamedQuery(`SELECT * FROM users WHERE id=:fn`, map[string]interface{}{"fn": u.ID})
	if err != nil {
		return err
	}
	if !row.Next() {
		return errors.New("failed to get user by id")
	}

	err = row.StructScan(&u)
	if err != nil {
		return err
	}

	return nil
}
