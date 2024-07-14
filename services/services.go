package services

import (
	"database/sql"
	"fmt"

	"github.com/gogineni1998/go-api/models"
	"github.com/gogineni1998/go-api/utilities"
)

func GetUsers(users *[]models.User, db *sql.DB) {
	var user models.User
	SELECT_QUERY := "SELECT * FROM users"
	rows, err := db.Query(SELECT_QUERY)
	utilities.ErrorHanler(err)
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Summary)
		utilities.ErrorHanler(err)
		*users = append(*users, user)
	}
}

func CreateUser(user *models.User, db *sql.DB) (int, error) {
	INSERT_QUERY := `INSERT INTO users (id, username, email, summary) VALUES($1, $2, $3, $4) RETURNING id`
	var id int
	err := db.QueryRow(INSERT_QUERY, user.ID, user.Username, user.Email, user.Summary).Scan(&id)
	utilities.ErrorHanler(err)
	if err != nil {
		return id, err
	}
	return id, nil
}

func UpdateUser(user *models.User, db *sql.DB) (int64, error) {
	fmt.Println(user)
	UPDATE_QUERY := "UPDATE users SET username=$2, email=$3, summary=$4 WHERE id=$1"
	row, err := db.Exec(UPDATE_QUERY, user.ID, user.Username, user.Email, user.Summary)
	utilities.ErrorHanler(err)
	return row.RowsAffected()
}

func DeleteUser(id int, db *sql.DB) (int64, error) {
	DELETE_QUERY := "DELETE FROM users where id=$1"
	_, err := db.Exec(DELETE_QUERY, id)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

func GetUser(id int, user *models.User, db *sql.DB) error {
	SELECT_QUERY := "SELECT * FROM users where id=$1"
	row := db.QueryRow(SELECT_QUERY, id)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Summary)
	utilities.ErrorHanler(err)
	return err
}
