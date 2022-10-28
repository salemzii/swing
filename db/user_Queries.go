package db

import (
	"context"
	"database/sql"
	"errors"
	"log"

	"github.com/salemzii/swing/users"
)

var (
	ErrIncorrectPassword = errors.New("incorrect password")
	ErrUserNotFound      = errors.New("no user found")
)

type CreateUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (repo SingleStoreRepository) CreateUser(user users.User) (*users.User, error) {
	res, err := repo.db.Exec(createUser, user.Username, user.Email, user.Password)

	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}
	user.Id = int(id)

	return &user, nil
}

func (repo SingleStoreRepository) LoginUser(logins users.LoginUser) (*users.User, error) {
	stmt, err := repo.db.Prepare(getUserByEmail)
	if err != nil {
		return &users.User{}, nil
	}

	defer stmt.Close()
	var user users.User

	row := stmt.QueryRowContext(context.Background(), logins.Email)
	err = row.Scan(&user.Id, &user.Email, &user.Password, &user.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			return &users.User{}, ErrUserNotFound
		}
		return &users.User{}, err
	}

	if user.Password != logins.Password {
		return &users.User{}, ErrIncorrectPassword
	}
	log.Println(user)
	return &user, nil
}
