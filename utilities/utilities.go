package utilities

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB_Connection *sql.DB

func init() {
	DB_Connection = EstablishConnection()
}

func EstablishConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://root:root@localhost:5432/test_db?sslmode=disable")
	if err != nil {
		log.Println(err)
	}
	err = db.Ping()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Connection Established Successfullty.....")
	return db
}

func CreateTable() string {
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INT PRIMARY KEY,
		username VARCHAR(50) NOT NULL UNIQUE,
		email VARCHAR(100) NOT NULL,
		summary VarCHAR(100) NOT NULL
	);`
	_, err := DB_Connection.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	return "Created Successfully"
}

func CloseConnection() {
	DB_Connection.Close()
}
