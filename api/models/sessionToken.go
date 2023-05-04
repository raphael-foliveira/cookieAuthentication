package models

import (
	"time"
)

type SessionToken struct {
	Id        string    `json:"id"`
	Token     string    `json:"token"`
	UserId    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}
