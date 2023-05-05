package database

import (
	"fmt"
	"time"

	"github.com/raphael-foliveira/cookieAuthentication/api/models"
)

func FindUser(id string) (models.User, error) {
	row := Db.QueryRow("SELECT id, username, password, email, created_at, deleted_at FROM users WHERE id = $1", id)
	var user models.User
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.DeletedAt)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func FindUserByUsername(username string) (models.User, error) {
	row := Db.QueryRow(`
	SELECT id, username, password, email, created_at, deleted_at 
	FROM users WHERE username = $1`,
		username)
	var user models.User
	err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.CreatedAt, &user.DeletedAt)
	if err != nil {
		fmt.Println(err.Error())
		return models.User{}, err
	}
	return user, nil
}

func CreateUser(username, password, email string) (models.User, error) {
	var user models.User
	row := Db.QueryRow(`
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
		return models.User{}, err
	}
	return user, nil
}

func FindUserBySessionToken(token string) (models.User, error) {
	var user models.User
	var sessionToken models.SessionToken
	row := Db.QueryRow(`
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
		return models.User{}, err
	}

	if time.Now().After(sessionToken.CreatedAt.Add(24 * 7 * time.Hour)) {
		fmt.Println("token expired. deleting...")
		_, err := Db.Exec("DELETE FROM session_tokens WHERE id = $1", sessionToken.Id)
		if err != nil {
			fmt.Println("error deleting session token:", err.Error())
		}
		return models.User{}, fmt.Errorf("token expired")
	}
	return user, nil
}
