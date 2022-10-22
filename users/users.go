package users

import "github.com/google/uuid"

type User struct {
	Id           uuid.UUID     `json:"id"`
	Username     string        `json:"username"`
	Password     string        `json:"password"`
	Email        string        `json:"email"`
	Applications []Application `json:"applications"`
}
