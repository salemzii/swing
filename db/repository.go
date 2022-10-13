package db

import "github.com/salemzii/swing/logs"

type Repository interface {
	Migrate() error
	Create(logrecord logs.LogRecord) (*logs.LogRecord, error)
	All() ([]logs.LogRecord, error)
	GetByLevel(string)
	GetByLineNum(int)
	GetByFunction(string)
	GetByOffset(int)
	Delete(id int64) error
}
