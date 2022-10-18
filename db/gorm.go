package db

import (
	"github.com/salemzii/swing/logs"
	"gorm.io/gorm"
)

type Record struct {
	gorm.Model
	Record logs.LogRecord `gorm:"embedded"`
}

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {

	return &GormRepository{
		db: db,
	}
}

func (repo GormRepository) GetByFunction(name string, limit ...int) (rcds []logs.LogRecord, err error) {
	return []logs.LogRecord{}, nil
}

func (repo GormRepository) GetByLevel(level string, limit ...int) (rcds []logs.LogRecord, err error) {
	return []logs.LogRecord{}, nil
}

func (repo GormRepository) GetByLineNum(line int, limit ...int) (rcds []logs.LogRecord, err error) {
	return []logs.LogRecord{}, nil
}

func (repo GormRepository) GetByOffset(offset int, limit ...int) (rcds []logs.LogRecord, err error) {
	return []logs.LogRecord{}, nil
}

func (repo GormRepository) All(limit ...int) (rcds []logs.LogRecord, err error) {
	return []logs.LogRecord{}, nil
}

func (repo GormRepository) GetById(id int) (rcd *logs.LogRecord, err error) {
	return &logs.LogRecord{}, nil
}

func (repo GormRepository) Create(logrecord logs.LogRecord) (*logs.LogRecord, error) {
	return &logs.LogRecord{}, nil
}
func (repo GormRepository) CreateMany(logrecords []logs.LogRecord) ([]logs.LogRecord, error) {
	return []logs.LogRecord{}, nil
}
func (repo GormRepository) Migrate() error {
	return nil
}
