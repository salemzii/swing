package users

import "time"

type User struct {
	Id       int       `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Created  time.Time `json:"created"`
}

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
