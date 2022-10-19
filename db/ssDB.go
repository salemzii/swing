package db

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"strconv"
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
				&record.Offset, &record.TimeStamp, &record.Created, &record.Logger)

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
				&record.Offset, &record.TimeStamp, &record.Created, &record.Logger)

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
				&record.Offset, &record.TimeStamp, &record.Created, &record.Logger)

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
				&record.Offset, &record.TimeStamp, &record.Created, &record.Logger)

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
				&record.Offset, &record.Created, &record.TimeStamp, &record.Logger)

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
		logrecord.Offset, logrecord.TimeStamp, logrecord.Logger)

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

func (repo SingleStoreRepository) Last15Minutes() (rcds []logs.LogRecord, err error) {
	stmt, err := repo.db.PrepareContext(context.Background(), getLast15Minutes)
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
				&record.Offset, &record.Created, &record.TimeStamp, &record.Logger)

			if err != nil {
				log.Fatal(err)
			}
			records = append(records, record)

		}
	}

	return records, nil
}

func (repo SingleStoreRepository) LastXMinutes(minutes int) (rcds []logs.LogRecord, err error) {
	stmt, err := repo.db.PrepareContext(context.Background(), getLastXMinutes)
	if err != nil {
		return
	}
	defer stmt.Close()
	var records []logs.LogRecord

	rows, err := stmt.QueryContext(context.Background(), minutes)
	if err != nil {
		return
	}
	{

		for rows.Next() {
			var record logs.LogRecord
			err := rows.Scan(&record.Id, &record.Level, &record.Message,
				&record.StackTrace, &record.Function, &record.LineNumber,
				&record.Offset, &record.Created, &record.TimeStamp, &record.Logger)

			if err != nil {
				log.Fatal(err)
			}
			records = append(records, record)

		}
	}

	return records, nil
}

func (repo SingleStoreRepository) CreateMany(logrecords []logs.LogRecord) (uint, error) {
	query := insertMany
	var inserts []string
	var params []interface{}

	for _, v := range logrecords {
		inserts = append(inserts, "(?, ?, ?, ?, ?, ?, ?, ?)")
		params = append(params, v.Message, v.Level, v.StackTrace, v.Function, v.LineNumber, v.Offset, v.TimeStamp, v.Logger)
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
			return 0, err
		}
		defer stmt.Close()
		res, err := stmt.ExecContext(ctx, params...)
		if err != nil {
			log.Printf("Error %s when inserting row into products table", err)
			return 0, err
		}
		rows, err := res.RowsAffected()
		if err != nil {
			log.Printf("Error %s when finding rows affected", err)
			return 0, err
		}
		log.Printf("%d records created simultaneously", rows)
		return uint(rows), err
	}
}

type DeleteRecord struct {
	Id uint64 `json:"id"`
}

type DeleteRecords struct {
	Ids []DeleteRecord `json:"ids"`
}

func (repo SingleStoreRepository) DeleteById(id uint64) (int, error) {
	stmt, err := repo.db.PrepareContext(context.Background(), delete)

	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(context.Background(), id)
	if err != nil {
		log.Printf("Error %s when Querying SQL ", err)
		return 0, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return 0, err
	}
	if rows != 0 {
		return int(rows), nil
	}
	return 0, errors.New("zero rows affected")
}

func (repo SingleStoreRepository) DeleteManyById(id []DeleteRecord) (rowsaffected int64, err error) {
	ids := []string{}
	for _, v := range id {
		ids = append(ids, strconv.Itoa(int(v.Id)))
	}

	values := "(" + strings.Join(ids, ",") + ")"
	log.Println(values)

	query := deleteMany + values
	log.Println(query)

	stmt, err := repo.db.PrepareContext(context.Background(), query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.ExecContext(context.Background())
	if err != nil {
		log.Printf("Error %s when Querying SQL ", err)
		return 0, err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
		return 0, err
	}
	if rows != 0 {
		return rows, nil
	}
	return 0, errors.New("zero rows affected")
}

func (repo SingleStoreRepository) Migrate() error {

	_, err := repo.db.Exec(migrate)
	if err != nil {
		return err
	}
	return nil
}
