package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

//Contectar() - Abre a conexão com o database
func Conectar() (*sql.DB, error){
	
	//"usuario:senha@/databaseName"
	stringConn := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConn)
	if err != nil {
		return nil, err
	}
	
	//testando conexão
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}