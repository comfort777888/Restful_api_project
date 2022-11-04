package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"rest_api_project/pkg/users"

	"github.com/gorilla/mux"
)

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Fatalf("error in getting id in handlers:%v", err)
	}
	u := &users.User{
		ID: id,
	}

	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("error in reading request: %v", err)
	}

	err = json.Unmarshal(bb, u)
	if err != nil {
		log.Printf("error in Unmarshal:%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = u.Create()
	if err != nil {
		log.Printf("error in create user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("user was created")
	w.WriteHeader(http.StatusCreated)
	fmt.Println(u)
}

func GetUserHandlerById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Fatalf("error in getting id in handlers:%v", err)
	}
	u := &users.User{
		ID: id,
	}
	err = u.Read(id)
	if err != nil {
		log.Fatalf("error in reading request: %v", err)
	}

	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		fmt.Println("error when encoding", "error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	u := &users.User{}
	uu, err := u.All()
	if err != nil {
		log.Fatalf("error in reading all users: %v", err)
	}
	err = json.NewEncoder(w).Encode(uu)
	if err != nil {
		fmt.Println("error when encoding in all users", "error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Fatalf("error in getting id in handlers:%v", err)
	}
	u := &users.User{
		ID: id,
	}

	bb, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("error in reading request: %v", err)
	}

	err = json.Unmarshal(bb, u)
	if err != nil {
		log.Printf("error in Unmarshal:%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = u.Update(id, u.Data)
	if err != nil {
		log.Printf("error in update user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("user was updated")
	w.WriteHeader(http.StatusOK)
	fmt.Println(u)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Fatalf("error in getting id in handlers:%v", err)
	}
	u := &users.User{
		ID: id,
	}
	err = u.Delete(id)
	if err != nil {
		log.Printf("error in delete user: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println("user was deleted")
	w.WriteHeader(http.StatusNoContent)
}
