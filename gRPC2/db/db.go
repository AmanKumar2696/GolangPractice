package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//To connect to database

func Conn() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:aman@123@tcp(127.0.0.1:3306)/WebApp")
	if err != nil {
		panic(err)
	}
	return db
}
