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
	HostName           = os.Getenv("singlestoreDbHost")
	Port               = os.Getenv("singlestorePort")
	USERNAME           = os.Getenv("singlestoreUsername")
	PASSWORD           = os.Getenv("singlestorePassword")
	DATABASE           = os.Getenv("singlestoreDatatbase")
	ErrCannotConnectDb = errors.New("unable to connect to database")
	ErrDuplicate       = errors.New("record already exists")
	ErrNotExists       = errors.New("row not exists")
	ErrDeleteFailed    = errors.New("delete failed")

	swingRepository *db.SingleStoreRepository
)

func init() {
	connection := USERNAME + ":" + PASSWORD + "@tcp(" + HostName + ":" + Port + ")/" + DATABASE + "?parseTime=true"
	fmt.Println(connection)

	database, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	err = database.Ping()
	if err != nil {
		panic(err)
	}

	swingRepository = db.NewSingleStoreRepository(database)
	if err := swingRepository.Migrate(); err != nil {
		log.Fatal(err)
	}
}

type AllRecordStruct struct {
	Limit int `json:"limit"`
}

func AllRecords(ctx context.Context, arg *AllRecordStruct) (rcds []logs.LogRecord, err error) {
	if arg.Limit == 0 {
		return nil, errors.New("limit cannot be 0")
	}
	records, err := swingRepository.All(arg.Limit)
	if err != nil {
		log.Println(err)
		return []logs.LogRecord{}, err
	}
	return records, nil

}
