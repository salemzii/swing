package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/salemzii/swing/db"
	"github.com/salemzii/swing/logs"
)

var (
	HostName = os.Getenv("singlestoreDbHost")
	Port     = os.Getenv("singlestorePort")
	USERNAME = os.Getenv("singlestoreUsername")
	PASSWORD = os.Getenv("singlestorePassword")
	DATABASE = os.Getenv("singlestoreDatabase")

	ErrCannotConnectDb = errors.New("unable to connect to database")
	ErrDuplicate       = errors.New("record already exists")
	ErrNotExists       = errors.New("row not exists")
	ErrDeleteFailed    = errors.New("delete failed")

	swingRepository *db.SingleStoreRepository
	gormRepository  *db.GormRepository
)

func init() {

	connection := USERNAME + ":" + PASSWORD + "@tcp(" + HostName + ":" + Port + ")/" + DATABASE + "?parseTime=true"
	fmt.Println(connection)

	database, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping()
	if err != nil {
		panic(err)
	}

	swingRepository = db.NewSingleStoreRepository(database)
	if err := swingRepository.Migrate(); err != nil {
		log.Fatal(err)
	}

	/*
		connection := USERNAME + ":" + PASSWORD + "@tcp(" + HostName + ":" + Port + ")/" + DATABASE + "?parseTime=True&loc=Local"

		gormDb, err := gorm.Open(gormMySQL.Open(connection), &gorm.Config{CreateBatchSize: 1000})
		if err != nil {
			log.Fatal(err)
		}

		gormRepository = db.NewGormRepository(gormDb)
		if err = gormRepository.Migrate(); err != nil {
			log.Fatal(err)
		}
	*/
}

type AllRecordStruct struct {
	Limit int `json:"limit"`
}
type RecordLineNum struct {
	Limit int `json:"limit"`
	Line  int `json:"line"`
}
type RecordLevel struct {
	Level string `json:"level"`
}
type RecordFunction struct {
	Function string `json:"function"`
}
type Record struct {
	Records []logs.LogRecord `json:"records"`
}

func CreateRecord(ctx context.Context, arg *logs.LogRecord) (*logs.LogRecord, error) {
	createdRecord, err := swingRepository.Create(*arg)
	if err != nil {
		return &logs.LogRecord{}, err
	}

	return createdRecord, nil
}

func CreateRecords(ctx context.Context, arg *Record) (uint, error) {
	rows, err := swingRepository.CreateMany(*&arg.Records)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func AllRecords(ctx context.Context, arg *AllRecordStruct) (rcds []logs.LogRecord, err error) {
	if arg.Limit == 0 {
		return nil, errors.New("limit cannot be 0")
	}
	records, err := swingRepository.All(arg.Limit)
	if err != nil {
		log.Println("ERROR", err)
		return []logs.LogRecord{}, err
	}
	return records, nil

}

func GetLast15MinutesRecords(ctx context.Context) (rcd []logs.LogRecord, err error) {

	records, err := swingRepository.Last15Minutes()
	if err != nil {
		log.Println("ERROR", err)
		return []logs.LogRecord{}, err
	}
	return records, nil
}

func GetRecordByNum(ctx context.Context, arg *RecordLineNum) (rcd []logs.LogRecord, err error) {

	record, err := swingRepository.GetByLineNum(arg.Line)
	if err != nil {
		log.Println("ERROR", err)
		return []logs.LogRecord{}, err
	}

	return record, nil
}

func GetRecordByLevel(ctx context.Context, arg *RecordLevel) (rcd []logs.LogRecord, err error) {
	record, err := swingRepository.GetByLevel(arg.Level)
	if err != nil {
		log.Println("ERROR", err)
		return []logs.LogRecord{}, err
	}

	return record, nil
}

func GetRecordByFunction(ctx context.Context, arg *RecordFunction) (rcd []logs.LogRecord, err error) {
	record, err := swingRepository.GetByFunction(arg.Function)
	if err != nil {
		log.Println("ERROR", err)
		return []logs.LogRecord{}, err
	}

	return record, nil
}
