package main

import (
	"fmt"
	"log"
	"net/http"
	
	"github.com/psdiaspedro/CRUD-go/server"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/users", server.CreateUser).Methods("POST")
	router.HandleFunc("/users", server.GetAllUsers).Methods("GET")
	router.HandleFunc("/users/{id}", server.GetUser).Methods("GET")
	router.HandleFunc("/users/{id}", server.UpdateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", server.DeleteUser).Methods("DELETE")

	fmt.Println("Listening on port 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}