package db

import (
	"database/sql"
	"log"

	"context"
)

const DefaultLimit = 100

type SingleStoreRepository struct {
	db *sql.DB
}

func NewSingleStoreRepository(db *sql.DB) *SingleStoreRepository {

	return &SingleStoreRepository{
		db: db,
	}
}
func (repo SingleStoreRepository) Migrate() error {

	log.Println("Migrating Table records")
	_, err := repo.db.Exec(migrate)
	if err != nil {
		return err
	}
	log.Println("Finished migrating Table records")

	log.Println("Migrating Table users")

	_, err = repo.db.ExecContext(context.Background(), migrateUser)
	if err != nil {
		return err
	}
	log.Println("Finished migrating Table users")
	log.Println("Migrating Table tokens")
	_, err = repo.db.ExecContext(context.Background(), migrateToken)
	if err != nil {
		return err
	}
	log.Println("Finished migrating Table tokens")

	return nil
}
