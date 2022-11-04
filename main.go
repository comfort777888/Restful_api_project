package main

import (
	"net/http"
	"rest_api_project/pkg/db"
	"rest_api_project/web/user"

	"github.com/gorilla/mux"
)

func main() {

	defer db.DB.Close()
	router := mux.NewRouter()
	router.HandleFunc("/user/{id}", user.CreateUserHandler).Methods(http.MethodPost)
	router.HandleFunc("/user/{id}", user.GetUserHandlerById).Methods(http.MethodGet)
	router.HandleFunc("/user/{id}", user.DeleteUserHandler).Methods(http.MethodDelete)
	router.HandleFunc("/user/{id}", user.UpdateUserHandler).Methods(http.MethodPut)
	router.HandleFunc("/users", user.GetAllUsersHandler).Methods(http.MethodGet)
	http.ListenAndServe(":8080", router)
}
