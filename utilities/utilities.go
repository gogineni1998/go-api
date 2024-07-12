package utilities

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ErrorHanler(err error) {
	if err != nil {
		log.Println(err)
	}
}
func EstablishConnection() *sql.DB {
	db, err := sql.Open("postgres", "postgres://root:root@localhost:5432/test_db?sslmode=disable")
	ErrorHanler(err)
	ErrorHanler(db.Ping())
	fmt.Println("Connection Established Successfullty.....")
	return db
}

func CreateTable() {
	db := EstablishConnection()
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS users (
		id INT PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		email VARCHAR(100) NOT NULL
	);`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

}