package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func Conect() (*sql.DB, error){
	
	stringConn := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", stringConn)
	if err != nil {
		return nil, err
	}
	
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}