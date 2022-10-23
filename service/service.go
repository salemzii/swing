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
	"github.com/salemzii/swing/users"
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

	SwingRepository *db.SingleStoreRepository
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

	database.SetMaxIdleConns(20)
	SwingRepository = db.NewSingleStoreRepository(database)
	if err := SwingRepository.Migrate(); err != nil {
		log.Fatal(err)
	}

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

type XRecords struct {
	Minutes int `json:"minutes"`
}

func CreateRecord(ctx context.Context, arg *logs.LogRecord) (*logs.LogRecord, error) {
	log.Println(arg)
	createdRecord, err := SwingRepository.Create(*arg)
	if err != nil {
		return &logs.LogRecord{}, err
	}

	return createdRecord, nil
}

func CreateRecords(ctx context.Context, arg *Record) (uint, error) {

	for _, v := range arg.Records {
		log.Println(v)
		switch v.Level {
		case "ERROR":
			// SEND ERROR MESSAGE To USER
			// Record TO Error Analytics DB
			//
		}
	}

	rows, err := SwingRepository.CreateMany(*&arg.Records)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func DeleteRecordF(ctx context.Context, arg *db.DeleteRecord) (int, error) {
	rows, err := SwingRepository.DeleteById(arg.Id)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func DeleteRecordsF(ctx context.Context, arg *db.DeleteRecords) (int64, error) {
	rows, err := SwingRepository.DeleteManyById(arg.Ids)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func AllRecords(ctx context.Context, arg *AllRecordStruct) (rcds []logs.LogRecord, err error) {
	if arg.Limit == 0 {
		return nil, errors.New("limit cannot be 0")
	}
	records, err := SwingRepository.All(arg.Limit)
	if err != nil {
		log.Println("ERROR", err)
		return []logs.LogRecord{}, err
	}
	return records, nil

}

func GetLast15MinutesRecords(ctx context.Context) (rcd []logs.LogRecord, err error) {

	records, err := SwingRepository.Last15Minutes()
	if err != nil {
		log.Println("ERROR", err)
		return []logs.LogRecord{}, err
	}
	return records, nil
}

func GetLastXMinutesRecords(ctx context.Context, arg *XRecords) (rcd []logs.LogRecord, err error) {

	records, err := SwingRepository.LastXMinutes(arg.Minutes)
	if err != nil {
		log.Println("ERROR", err)
		return []logs.LogRecord{}, err
	}
	return records, nil
}

func GetRecordByNum(ctx context.Context, arg *RecordLineNum) (rcd []logs.LogRecord, err error) {

	record, err := SwingRepository.GetByLineNum(arg.Line)
	if err != nil {
		log.Println("ERROR", err)
		return []logs.LogRecord{}, err
	}

	return record, nil
}

func GetRecordByLevel(ctx context.Context, arg *RecordLevel) (rcd []logs.LogRecord, err error) {
	record, err := SwingRepository.GetByLevel(arg.Level)
	if err != nil {
		log.Println("ERROR", err)
		return []logs.LogRecord{}, err
	}

	return record, nil
}

func GetRecordByFunction(ctx context.Context, arg *RecordFunction) (rcd []logs.LogRecord, err error) {
	record, err := SwingRepository.GetByFunction(arg.Function)
	if err != nil {
		log.Println("ERROR", err)
		return []logs.LogRecord{}, err
	}

	return record, nil
}

func VerifyToken(token string) (users.TokenDetails, error) {
	details, err := SwingRepository.FetchToken(token)
	if err != nil {
		log.Println("ERROR", err)
		return users.TokenDetails{}, err
	}
	return details, nil
}
