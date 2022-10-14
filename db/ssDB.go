package db

import (
	"database/sql"

	"github.com/salemzii/swing/logs"
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

func (repo SingleStoreRepository) GetByFunction(string) {

}

func (repo SingleStoreRepository) GetByLevel(string) {}

func (repo SingleStoreRepository) GetByLineNum(int) {}

func (repo SingleStoreRepository) GetByOffset(int) {}
func (repo SingleStoreRepository) All()            {}
func (repo SingleStoreRepository) Create(logrecord logs.LogRecord) (*logs.LogRecord, error) {
	return &logs.LogRecord{}, nil
}
func (repo SingleStoreRepository) Migrate() error {

	return nil
}
