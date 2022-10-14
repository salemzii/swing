package db

import (
	"database/sql"
	"log"

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

func (repo SingleStoreRepository) GetByFunction(name string, limit ...int) (rcds []logs.LogRecord, err error) {
	stmt, err := repo.db.Prepare(getByFunction)
	if err != nil {
		return
	}
	defer stmt.Close()

	var lim int
	var records []logs.LogRecord

	if len(limit) > 0 {
		lim = limit[0]
	}

	rows, err := stmt.Query(name, lim)
	if err != nil {
		return
	}

	{

		for rows.Next() {
			var record logs.LogRecord
			err := rows.Scan(&record.Id, &record.Level, &record.Message,
				&record.StackTrace, &record.Function, &record.LineNumber,
				&record.Offset, &record.TimeStamp, &record.Created)

			if err != nil {
				log.Fatal(err)
			}
			records = append(records, record)

		}
	}

	return records, nil
}

func (repo SingleStoreRepository) GetByLevel(level string, limit ...int) (rcds []logs.LogRecord, err error) {
	stmt, err := repo.db.Prepare(getByLevel)
	if err != nil {
		return
	}
	defer stmt.Close()

	var lim int
	var records []logs.LogRecord

	if len(limit) > 0 {
		lim = limit[0]
	}

	rows, err := stmt.Query(level, lim)
	if err != nil {
		return
	}

	{

		for rows.Next() {
			var record logs.LogRecord
			err := rows.Scan(&record.Id, &record.Level, &record.Message,
				&record.StackTrace, &record.Function, &record.LineNumber,
				&record.Offset, &record.TimeStamp, &record.Created)

			if err != nil {
				log.Fatal(err)
			}
			records = append(records, record)

		}
	}

	return records, nil
}

func (repo SingleStoreRepository) GetByLineNum(line int, limit ...int) (rcds []logs.LogRecord, err error) {
	stmt, err := repo.db.Prepare(getByLineNum)
	if err != nil {
		return
	}
	defer stmt.Close()

	var lim int
	var records []logs.LogRecord

	if len(limit) > 0 {
		lim = limit[0]
	}

	rows, err := stmt.Query(line, lim)
	if err != nil {
		return
	}

	{

		for rows.Next() {
			var record logs.LogRecord
			err := rows.Scan(&record.Id, &record.Level, &record.Message,
				&record.StackTrace, &record.Function, &record.LineNumber,
				&record.Offset, &record.TimeStamp, &record.Created)

			if err != nil {
				log.Fatal(err)
			}
			records = append(records, record)

		}
	}

	return records, nil
}

func (repo SingleStoreRepository) GetByOffset(offset int, limit ...int) (rcds []logs.LogRecord, err error) {
	stmt, err := repo.db.Prepare(getByLineNum)
	if err != nil {
		return
	}
	defer stmt.Close()

	var lim int
	var records []logs.LogRecord

	if len(limit) > 0 {
		lim = limit[0]
	}

	rows, err := stmt.Query(offset, lim)
	if err != nil {
		return
	}

	{

		for rows.Next() {
			var record logs.LogRecord
			err := rows.Scan(&record.Id, &record.Level, &record.Message,
				&record.StackTrace, &record.Function, &record.LineNumber,
				&record.Offset, &record.TimeStamp, &record.Created)

			if err != nil {
				log.Fatal(err)
			}
			records = append(records, record)

		}
	}

	return records, nil

}

func (repo SingleStoreRepository) All(limit ...int) (rcds []logs.LogRecord, err error) {
	stmt, err := repo.db.Prepare(all)
	if err != nil {
		return
	}
	defer stmt.Close()

	lim := 100
	var records []logs.LogRecord

	if len(limit) > 0 {
		lim = limit[0]
	}

	rows, err := stmt.Query(lim)
	if err != nil {
		return
	}
	{

		for rows.Next() {
			var record logs.LogRecord
			err := rows.Scan(&record.Id, &record.Level, &record.Message,
				&record.StackTrace, &record.Function, &record.LineNumber,
				&record.Offset, &record.TimeStamp, &record.Created)

			if err != nil {
				log.Fatal(err)
			}
			records = append(records, record)

		}
	}

	return records, nil

}
func (repo SingleStoreRepository) Create(logrecord logs.LogRecord) (*logs.LogRecord, error) {
	res, err := repo.db.Exec(create, logrecord.Message, logrecord.Level,
		logrecord.StackTrace, logrecord.Function, logrecord.LineNumber,
		logrecord.Offset, logrecord.TimeStamp)

	if err != nil {
		return nil, err
	}

	_, err = res.LastInsertId()

	if err != nil {
		return nil, err
	}

	return &logrecord, nil
}
func (repo SingleStoreRepository) Migrate() error {

	_, err := repo.db.Exec(migrate)
	if err != nil {
		return err
	}
	return nil
}
