package auth

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/raphael-foliveira/cookieAuthentication/api/database"
	"github.com/raphael-foliveira/cookieAuthentication/api/models"
)

func RenewSessionToken(userId int) (models.SessionToken, error) {
	rows, err := database.Db.Query("SELECT s.id FROM session_tokens s JOIN users u on s.user_id = u.id WHERE u.id = $1", userId)
	if err != nil {
		fmt.Println("error renewing session token:", err.Error())
		return models.SessionToken{}, err
	}
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			fmt.Println("error renewing session token:", err.Error())
			return models.SessionToken{}, err
		}
		_, err = database.Db.Exec("DELETE FROM session_tokens WHERE id = $1", id)
		if err != nil {
			fmt.Println("error renewing session token:", err.Error())
			return models.SessionToken{}, err
		}
	}
	return CreateSessionToken(userId)

}

func CreateSessionToken(userId int) (models.SessionToken, error) {
	uuid := uuid.NewString()
	row := database.Db.QueryRow("INSERT INTO session_tokens (token, user_id) VALUES ($1, $2) RETURNING id, token, user_id", uuid, userId)
	var sessionToken models.SessionToken
	err := row.Scan(&sessionToken.Id, &sessionToken.Token, &sessionToken.UserId)
	if err != nil {
		fmt.Println("error creating sesssion token:", err.Error())
		return models.SessionToken{}, err
	}
	return sessionToken, nil
}

func ValidateSessionToken(token string) (models.SessionToken, error) {
	row := database.Db.QueryRow("SELECT id, token, user_id FROM session_tokens WHERE token = $1", token)
	var sessionToken models.SessionToken
	err := row.Scan(&sessionToken.Id, &sessionToken.Token, &sessionToken.UserId)
	if err != nil {
		fmt.Println("error validating sesssion token:", err.Error())
		return models.SessionToken{}, err
	}
	return sessionToken, nil
}
