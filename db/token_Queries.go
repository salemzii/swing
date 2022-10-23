package db

import (
	"context"

	"github.com/salemzii/swing/users"
)

func (repo SingleStoreRepository) FetchToken(token string) (users.TokenDetails, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), GetToken)
	if err != nil {
		return users.TokenDetails{}, err
	}
	defer stmt.Close()
	var details users.TokenDetails
	row, err := stmt.QueryContext(context.Background(), token)
	if err != nil {
		return users.TokenDetails{}, err
	}
	err = row.Scan(&details.Id, &details.Token, &details.Expires_at, &details.Rate_limit, &details.Enabled, &details.Created)
	if err != nil {
		return users.TokenDetails{}, err
	}
	return details, nil
}
