package services

import (
	"database/sql"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/gogineni1998/go-api/models"
)

var psql = sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

func GetUsers(users *[]models.User, db *sql.DB) {
	var user models.User
	SELECT_QUERY := psql.Select("*").From("users")
	rows, err := SELECT_QUERY.RunWith(db).Query()
	if err != nil {
		log.Println(err)
	}
	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Summary)
		if err != nil {
			log.Println(err)
		}
		*users = append(*users, user)
	}
}

func CreateUser(user *models.User, db *sql.DB) (int, error) {
	INSERT_QUERY := psql.Insert("users").Columns("id", "username", "email", "summary").Values(user.ID, user.Username, user.Email, user.Summary).Suffix("RETURNING id")
	var id int
	err := INSERT_QUERY.RunWith(db).QueryRow().Scan(&id)
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		return id, err
	}
	return id, nil
}

func UpdateUser(user *models.User, db *sql.DB) (int64, error) {
	UPDATE_QUERY := psql.Update("users").Set("username", user.Username).Set("email", user.Email).Set("summary", user.Summary).Where(sq.Eq{"id": user.ID})
	result, err := UPDATE_QUERY.RunWith(db).Exec()
	if err != nil {
		log.Println(err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
	}
	return rowsAffected, nil
}

func DeleteUser(id int, db *sql.DB) (int64, error) {
	DELETE_QUERY := psql.Delete("users").Where(sq.Eq{"id": id})
	result, err := DELETE_QUERY.RunWith(db).Exec()
	if err != nil {
		return 0, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}

func GetUser(id int, user *models.User, db *sql.DB) error {
	SELECT_QUERY := psql.Select("*").From("users").Where(sq.Eq{"id": id})
	row := SELECT_QUERY.RunWith(db).QueryRow()
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Summary)
	if err != nil {
		log.Println(err)
	}
	return err
}
