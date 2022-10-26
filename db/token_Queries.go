package db

import (
	"context"

	"github.com/salemzii/swing/users"
)

func (repo SingleStoreRepository) CreateToken(token users.TokenDetails) (tk users.TokenDetails, err error) {
	res, err := repo.db.ExecContext(context.Background(), createToken, token.Token,
		token.Expires_at, token.Rate_limit, token.Enabled, token.UserId)
	if err != nil {
		return users.TokenDetails{}, err
	}
	id, err := res.LastInsertId()

	if err != nil {
		return users.TokenDetails{}, err
	}
	token.Id = int(id)
	return token, nil
}

func (repo SingleStoreRepository) FetchToken(token string) (users.TokenDetails, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), getToken)
	if err != nil {
		return users.TokenDetails{}, err
	}
	defer stmt.Close()
	var details users.TokenDetails
	row := stmt.QueryRowContext(context.Background(), token)
	err = row.Scan(&details.Id, &details.Token, &details.Expires_at, &details.Rate_limit, &details.Enabled, &details.Created)
	if err != nil {
		return users.TokenDetails{}, err
	}
	return details, nil
}
