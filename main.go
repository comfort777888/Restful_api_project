package main

import (
	"cccccccc/pkg/db"
	"cccccccc/pkg/users"
	"fmt"
	"log"
)

func main() {
	s := users.NewStore(db.Newdb())
	defer s.DB.Close()

	u := users.NewUser(1, "Aray")
	err := u.Create(s)
	if err != nil {
		log.Fatalf("error in creating user %v", err)
	}

	fmt.Println("До изменения и удаления:")

	err = u.Read(s)
	if err != nil {
		log.Fatalf("error in reading User %v", err)
	}

	err = u.Update(s, 2, "Bill")
	if err != nil {
		log.Fatalf("error in updating User %v", err)
	}

	err = u.Delete(s, 3)
	if err != nil {
		log.Fatalf("error in deleting User %v", err)
	}

	fmt.Println("После изменения и удаления:")

	err = u.Read(s)
	if err != nil {
		log.Fatalf("error in reading User after updating and deleting %v", err)
	}

	err = u.ReadRow(s, 2)
	if err != nil {
		log.Fatalf("error in reading selected row %v", err)
	}
}
