package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

const user = "root"
const password = "1234"
const host = "localhost"
const port = 3306
const database = "crud"

func OpenDB() {
	conn, err := sql.Open("mysql", generateURL())
	if err != nil {
		panic(err)
	}
	db = conn
}

func Ping() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}

func CloseDB() {
	db.Close()
}

func generateURL() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, database)
}
