package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/raphael-foliveira/cookieAuthentication/api/database"
)

type User struct {
	Id        string       `json:"id"`
	Username  string       `json:"username"`
	Password  string       `json:"password"`
	Email     string       `json:"email"`
	CreatedAt time.Time    `json:"created_at"`
	DeletedAt sql.NullTime `json:"deleted_at"`
}

func FindUser(id string) (User, error) {
	row := database.Db.QueryRow("SELECT id, username, password, email, created_at, deleted_at FROM users WHERE id = $1", id)
	var user User
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.DeletedAt)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserByUsername(username string) (User, error) {
	row := database.Db.QueryRow("SELECT id, username, password, email, created_at, deleted_at FROM users WHERE username = $1", username)
	var user User
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.DeletedAt)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}
	return user, nil
}

func CreateUser(username, password, email string) (User, error) {
	var user User
	err := database.Db.QueryRow("INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id, username, password, email, created_at, deleted_at", username, password, email).Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.DeletedAt)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}
	return user, nil
}
