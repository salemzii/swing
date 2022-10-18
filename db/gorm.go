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
	var records []logs.LogRecord

	if len(limit) > 0 {
		if err := repo.db.Raw(all, limit[0]).Scan(&records).Error; err != nil {
			return []logs.LogRecord{}, err
		}
		return records, nil
	}

	if err := repo.db.Raw(all, 500).Scan(&records).Error; err != nil {
		return []logs.LogRecord{}, err
	}
	return records, nil
}

func (repo GormRepository) GetById(id int) (rcd *logs.LogRecord, err error) {
	return &logs.LogRecord{}, nil
}

func (repo GormRepository) Create(logrecord Record) (*Record, error) {

	tx := repo.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}

	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&logrecord).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return &logrecord, nil
}

func (repo GormRepository) CreateMany(logrecords []Record) ([]Record, error) {

	tx := repo.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}

	}()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.CreateInBatches(&logrecords, 100).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}
	return logrecords, nil
}

func (repo GormRepository) Migrate() error {
	// Add table suffix when creating tables
	//repo.db.AutoMigrate(&Record{}, &users.User{})

	if err := repo.db.Raw(migrate).Error; err != nil {
		return err
	}
	return nil
}
