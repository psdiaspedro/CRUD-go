package main

import (
	//"database/sql"
	"fmt"
	"log"
	"net/http"
	"github.com/psdiaspedro/CRUD-go/server"
	
	"github.com/gorilla/mux"
)

func main() {
	
	//CRUD - Create, Read, Update, Delete

	//CREATE - POST
	//READ - GET
	//UPDATE - PUT
	//DELETE - DELETE

	router := mux.NewRouter()
	router.HandleFunc("/usuarios", server.CreateUser).Methods("POST")
	router.HandleFunc("/usuarios", server.GetAllUsers).Methods("GET")
	router.HandleFunc("/usuarios/{id}", server.GetUser).Methods("GET")
	router.HandleFunc("/usuarios/{id}", server.UpdateUser).Methods("PUT")
	router.HandleFunc("/usuarios/{id}", server.DeleteUser).Methods("DELETE")

	fmt.Println("Escutando na porta 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}