package db

import (
	"context"

	"github.com/salemzii/swing/users"
)

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
