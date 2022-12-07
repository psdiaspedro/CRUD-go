package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/psdiaspedro/CRUD-go/database"
)

type user struct {
	ID		uint32 `json:"id"`
	Name	string `json:"name"`
	Email	string `json:"email"`
}

//CreateUser() - Creates an user inside database
func CreateUser(w http.ResponseWriter, r *http.Request) {

	//passo 1 - ler o corpo da request (com os dados do usuário)
	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Reading body request has failed!"))
		return
	}

	//passo 2 - criar um usuário do tipo da struct (ja no formato que queremos)
	var user user

	//passo 3 - Parsear o income JSON para user struct
	if err = json.Unmarshal(bodyReq, &user); err != nil {
		w.Write([]byte("Parsing JSON to struct has failed!"))
		return
	}

	//passo 4 - Conectar no banco de dados
	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("Contecting to database has failed!"))
		return
	}
	defer db.Close()

	// PREPARE STATEMENT
	statement, err := db.Prepare("insert into usuarios (name, email) values(?, ?)")
	if err != nil {
		w.Write([]byte("creating statement has failed!"))
		return
	}
	defer statement.Close()


	insert, err := statement.Exec(user.Name, user.Email)
	if err != nil {
		w.Write([]byte("Executing statement has failed!"))
		return
	}

	idInserted, err := insert.LastInsertId()
	if err != nil {
		w.Write([]byte("get last ID has failed!"))
		return
	}

	//STATUS CODE
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("User has been created! Id: %d", idInserted)))

}

//GetUser() - Get all users from database
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	
	//Passo 1
	//Como não vem corpo (pois é um get), a primeira coisa é abrir a conexão com o db
	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("Contecting to database has failed!"))
		return
	}
	defer db.Close()

	//Passo 2 - Fazer a busca dos usuários
	// SELECT * FROM usuarios
	rows, err := db.Query("SELECT * from usuarios")
	if err != nil {
		w.Write([]byte("Get users has failed"))
		return
	}
	defer rows.Close()

	var users []user
	
	for rows.Next() {
		var user user
		
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("Scaning user has failed"))
			return
		}
		users = append(users, user)
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(users); err != nil {
		w.Write([]byte("Parsing user to JSON has failed"))
		return
	}
}

//GetUser() - Get an user from database
func GetUser(w  http.ResponseWriter, r *http.Request) {

	//Passo 1
	//Nesse caso tem um parâmetro na rota, primeiro passo é ler
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Convert id param to int has failed!"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("conect to database has failed!"))
		return
	}

	row, err := db.Query("select * from usuarios where id = ?", ID)
	if err != nil {
		w.Write([]byte("get user by ID has failed!"))
		return
	}

	var user user
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			w.Write([]byte("scan user has failed!"))
			return
		}
	}

	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.Write([]byte("Parse JSON to struct has failed!"))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Convert id param to int has failed!"))
		return
	}

	bodyReq, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Reading body request has failed!"))
		return
	}

	var user user
	if err := json.Unmarshal(bodyReq, &user); err != nil {
		w.Write([]byte("Parse JSON to struct has failed!"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("Contecting to database has failed!"))
		return
	}
	defer db.Close()

	// PREPARE STATEMENT
	statement, err := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if err != nil {
		w.Write([]byte("creating statement has failed!"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, ID); err != nil {
		w.Write([]byte("Executing statement has failed!"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	
	params := mux.Vars(r)

	ID, err := strconv.ParseUint(params["id"], 10, 32)
	if err != nil {
		w.Write([]byte("Convert id param to int has failed!"))
		return
	}

	db, err := database.Conectar()
	if err != nil {
		w.Write([]byte("Contecting to database has failed!"))
		return
	}
	defer db.Close()

	// PREPARE STATEMENT
	statement, err := db.Prepare("delete from usuarios where id = ?")
	if err != nil {
		w.Write([]byte("creating statement has failed!"))
		return
	}
	defer statement.Close()

	if _, err := statement.Exec(ID); err != nil {
		w.Write([]byte("deleting user has failed!"))
		return
	}
}
