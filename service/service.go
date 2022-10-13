package service

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/salemzii/swing/db"
)

var (
	HostName        = os.Getenv("singlestoreDbHost")
	Port            = os.Getenv("singlestorePort")
	USERNAME        = os.Getenv("singlestoreUsername")
	PASSWORD        = os.Getenv("singlestorePassword")
	DATABASE        = os.Getenv("singlestoreDatatbase")
	swingRepository *db.SingleStoreRepository
)

func init() {
	connection := USERNAME + ":" + PASSWORD + "@tcp(" + HostName + ":" + Port + ")/" + DATABASE + "?parseTime=true"
	database, err := sql.Open("mysql", connection)
	if err != nil {
		panic(err)
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
