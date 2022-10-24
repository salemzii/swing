package db

import (
	"context"
	"errors"

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
	res, err := repo.db.Exec(createUser, user.Email, user.Username, user.Password)

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

func (repo SingleStoreRepository) LoginUser(logins users.LoginUser) (*users.User, string, error) {
	stmt, err := repo.db.Prepare(getUserByEmail)
	if err != nil {
		return &users.User{}, "", nil
	}

	defer stmt.Close()
	var user users.User
	var usertoken users.TokenDetails

	row := stmt.QueryRowContext(context.Background(), logins.Email)
	err = row.Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return &users.User{}, "", err
	}

	if user.Password != logins.Password {
		return &users.User{}, "", ErrIncorrectPassword
	}

	tokenStmt, err := repo.db.PrepareContext(context.Background(), getTokenByUserId)
	if err != nil {
		return &users.User{}, "", nil
	}

	defer tokenStmt.Close()
	tokenrow := tokenStmt.QueryRowContext(context.Background(), user.Id)

	err = tokenrow.Scan(&usertoken.Id, &usertoken.Token, &usertoken.Expires_at, &usertoken.Rate_limit, &usertoken.Enabled, &usertoken.UserId)
	if err != nil {
		return &users.User{}, "", err
	}

	return &users.User{}, "", nil
}
