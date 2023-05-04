package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/raphael-foliveira/cookieAuthentication/api/database"
)

type User struct {
	Id        int          `json:"id"`
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
	row := database.Db.QueryRow(`
	SELECT id, username, password, email, created_at, deleted_at 
	FROM users WHERE username = $1`,
		username)
	var user User
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.DeletedAt)
	if err != nil {
		fmt.Println(err.Error())
		return User{}, err
	}
	return user, nil
}

func CreateUser(username, password, email string) (User, error) {
	var user User
	row := database.Db.QueryRow(`
			INSERT INTO users (username, password, email) 
			VALUES ($1, $2, $3) 
			RETURNING id, username, password, email, created_at, deleted_at`,
		username,
		password,
		email)

	err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.Email,
		&user.CreatedAt,
		&user.DeletedAt)
	if err != nil {
		fmt.Println(err.Error())
		return User{}, err
	}
	return user, nil
}

func FindUserBySessionToken(token string) (User, error) {
	var user User
	var sessionToken SessionToken
	row := database.Db.QueryRow(`
		SELECT 
		u.id, u.username, u.password, u.email, u.created_at, s.id, s.created_at 
		FROM users u INNER JOIN session_tokens s 
		ON u.id = s.user_id WHERE s.token = $1
	`,
		token,
	)
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &sessionToken.Id, &sessionToken.CreatedAt)
	if err != nil {
		fmt.Println(err)
		return User{}, err
	}

	if time.Now().After(sessionToken.CreatedAt.Add(24 * 7 * time.Hour)) {
		fmt.Println("token expired. deleting...")
		_, err := database.Db.Exec("DELETE FROM session_tokens WHERE id = $1", sessionToken.Id)
		if err != nil {
			fmt.Println("error deleting session token:", err.Error())
		}
		return User{}, fmt.Errorf("token expired")
	}
	return user, nil
}
