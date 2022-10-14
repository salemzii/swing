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

// 8636dd65a36da865428e088c40f78a01f60532c33fd1599fa095825507af1eb2 Swing API KEy
