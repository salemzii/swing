package db

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"time"

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

	var records []logs.LogRecord

	rows, err := stmt.Query(name)
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

	var records []logs.LogRecord

	rows, err := stmt.Query(level)
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
	stmt, err := repo.db.PrepareContext(context.Background(), getByLineNum)
	if err != nil {
		return
	}
	defer stmt.Close()

	var records []logs.LogRecord
	rows, err := stmt.QueryContext(context.Background(), line)
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

	var records []logs.LogRecord

	rows, err := stmt.Query(offset)
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
	stmt, err := repo.db.PrepareContext(context.Background(), all)
	if err != nil {
		return
	}
	defer stmt.Close()
	var records []logs.LogRecord

	rows, err := stmt.QueryContext(context.Background())
	if err != nil {
		return
	}
	{

		for rows.Next() {
			var record logs.LogRecord
			err := rows.Scan(&record.Id, &record.Level, &record.Message,
				&record.StackTrace, &record.Function, &record.LineNumber,
				&record.Offset, &record.Created, &record.TimeStamp)

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

	id, err := res.LastInsertId()

	if err != nil {
		return nil, err
	}
	log.Println(id)

	logrecord.Id = int(id)

	return &logrecord, nil
}

func (repo SingleStoreRepository) CreateMany(logrecords []logs.LogRecord) ([]logs.LogRecord, error) {
	query := insertMany
	var inserts []string
	var params []interface{}

	for _, v := range logrecords {
		inserts = append(inserts, "(?, ?, ?, ?, ?, ?, ?)")
		params = append(params, v.Message, v.Level, v.StackTrace, v.Function, v.LineNumber, v.Offset, v.TimeStamp)
	}
	queryVals := strings.Join(inserts, ",")
	query = query + queryVals
	log.Println("query is", query)

	{

		ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancelfunc()
		stmt, err := repo.db.PrepareContext(ctx, query)
		if err != nil {
			log.Printf("Error %s when preparing SQL statement", err)
			return []logs.LogRecord{}, err
		}
		defer stmt.Close()
		res, err := stmt.ExecContext(ctx, params...)
		if err != nil {
			log.Printf("Error %s when inserting row into products table", err)
			return []logs.LogRecord{}, err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error %s when finding rows affected", err)
			return []logs.LogRecord{}, err
		}
		log.Printf("%d products created simulatneously", rows)
		return []logs.LogRecord{}, err
	}
}

func (repo SingleStoreRepository) Migrate() error {

	_, err := repo.db.Exec(migrate)
	if err != nil {
		return err
	}
	return nil
}
