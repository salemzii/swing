package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	jwt "github.com/golang-jwt/jwt"
	"github.com/salemzii/swing/app"
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
	JwtSecretKey    = []byte(os.Getenv("JwtSecretKey"))
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
	Tokenid string `json:"token"`
	Limit   int    `json:"limit"`
}
type RecordLineNum struct {
	Tokenid string `json:"token"`
	Limit   int    `json:"limit"`
	Line    int    `json:"line"`
}
type RecordLevel struct {
	Tokenid string `json:"token"`
	Level   string `json:"level"`
}
type RecordFunction struct {
	Tokenid  string `json:"token"`
	Function string `json:"function"`
}
type Record struct {
	Tokenid string           `json:"token"`
	Records []logs.LogRecord `json:"records"`
}

type XRecords struct {
	Tokenid string `json:"token"`
	Minutes int    `json:"minutes"`
}

func CreateRecord(ctx context.Context, arg *logs.LogRecord) (*logs.LogRecord, error) {
	log.Println(arg)
	_, userid, _, err := app.ServiceParseToken(arg.TokenId)
	if err != nil {
		return nil, err
	}

	if userid == 0 {
		return nil, errors.New("invalid userid")
	}

	arg.UserId = userid
	createdRecord, err := SwingRepository.Create(*arg)
	if err != nil {
		return &logs.LogRecord{}, err
	}

	return createdRecord, nil
}

// check
func CreateRecords(ctx context.Context, arg *Record) (uint, error) {
	_, userid, _, err := app.ServiceParseToken(arg.Tokenid)
	if err != nil {
		return 0, err
	}
	if userid == 0 {
		return 0, errors.New("invalid userid")
	}
	logRecords := []logs.LogRecord{}
	for _, v := range arg.Records {
		log.Println(v)
		v.UserId = userid
		logRecords = append(logRecords, v)
		switch v.Level {
		case "ERROR":
			// SEND ERROR MESSAGE To USER
			// Record TO Error Analytics DB

		}
	}

	rows, err := SwingRepository.CreateMany(logRecords)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func DeleteRecordF(ctx context.Context, arg *db.DeleteRecord) (int, error) {
	_, userid, _, err := app.ServiceParseToken(arg.Tokenid)
	if err != nil {
		return 0, err
	}
	rows, err := SwingRepository.DeleteById(userid, arg.Id)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func DeleteRecordsF(ctx context.Context, arg *db.DeleteRecords) (int64, error) {
	_, userid, _, err := app.ServiceParseToken(arg.Tokenid)
	if err != nil {
		return 0, err
	}
	rows, err := SwingRepository.DeleteManyById(userid, arg.Ids)
	if err != nil {
		return 0, err
	}
	return rows, nil
}

func AllRecords(ctx context.Context, arg *AllRecordStruct) (rcds []logs.LogRecord, err error) {
	_, userid, _, err := app.ServiceParseToken(arg.Tokenid)
	if err != nil {
		return nil, err
	}
	if arg.Limit == 0 {
		return nil, errors.New("limit cannot be 0")
	}
	records, err := SwingRepository.All(userid, arg.Limit)
	if err != nil {
		log.Println("ERROR", err)
		return []logs.LogRecord{}, err
	}
	return records, nil

}

func GetLast15MinutesRecords(ctx context.Context, arg *XRecords) (rcd []logs.LogRecord, err error) {
	_, userid, _, err := app.ServiceParseToken(arg.Tokenid)
	if err != nil {
		return []logs.LogRecord{}, err
	}
	records, err := SwingRepository.Last15Minutes(userid)
	if err != nil {
		log.Println("ERROR", err)
		return []logs.LogRecord{}, err
	}
	return records, nil
}

func GetLastXMinutesRecords(ctx context.Context, arg *XRecords) (rcd []logs.LogRecord, err error) {
	_, userid, _, err := app.ServiceParseToken(arg.Tokenid)
	if err != nil {
		return []logs.LogRecord{}, err
	}
	log.Println(arg)
	records, err := SwingRepository.LastXMinutes(userid, arg.Minutes)
	if err != nil {
		log.Println("ERROR", err)
		return []logs.LogRecord{}, err
	}
	return records, nil
}

func GetRecordByNum(ctx context.Context, arg *RecordLineNum) (rcd []logs.LogRecord, err error) {
	_, userid, _, err := app.ServiceParseToken(arg.Tokenid)
	if err != nil {
		return []logs.LogRecord{}, err
	}
	record, err := SwingRepository.GetByLineNum(arg.Line, userid)
	if err != nil {
		log.Println("ERROR", err)
		return []logs.LogRecord{}, err
	}

	return record, nil
}

func GetRecordByLevel(ctx context.Context, arg *RecordLevel) (rcd []logs.LogRecord, err error) {
	_, userid, _, err := app.ServiceParseToken(arg.Tokenid)
	if err != nil {
		return []logs.LogRecord{}, err
	}
	record, err := SwingRepository.GetByLevel(arg.Level, userid)
	if err != nil {
		log.Println("ERROR", err)
		return []logs.LogRecord{}, err
	}

	return record, nil
}

func GetRecordByFunction(ctx context.Context, arg *RecordFunction) (rcd []logs.LogRecord, err error) {
	_, userid, _, err := app.ServiceParseToken(arg.Tokenid)
	if err != nil {
		return []logs.LogRecord{}, err
	}
	record, err := SwingRepository.GetByFunction(arg.Function, userid)
	if err != nil {
		log.Println("ERROR", err)
		return []logs.LogRecord{}, err
	}

	return record, nil
}

func CreateUserAccount(ctx context.Context, arg *users.User) (*LoginResponse, error) {
	user, err := SwingRepository.CreateUser(*arg)

	if err != nil {
		return nil, err
	}

	token, err := GenerateJwt(time.Duration(44_640), user.Username, user.Id)
	if err != nil {
		return nil, err
	}

	var tk users.TokenDetails
	tk.Token = token
	tk.Rate_limit = 500
	tk.Expires_at = time.Now().Add(time.Duration(44_640) * time.Minute)
	tk.Enabled = true
	tk.UserId = user.Id

	t, err := SwingRepository.CreateToken(tk)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{User: *user, Token: t.Token}, nil
}

type LoginResponse struct {
	User  users.User
	Token string
}

func LoginUserAccount(ctx context.Context, arg *users.LoginUser) (*LoginResponse, error) {

	var resp LoginResponse
	user, err := SwingRepository.LoginUser(*arg)
	if err != nil {
		return nil, err
	}

	if user.Username == "" || user.Email == "" {
		return nil, errors.New("incorrect login credentials")
	}

	token, err := GenerateJwt(time.Duration(15), user.Username, user.Id)
	if err != nil {
		return nil, err
	}
	resp.User = *user
	resp.Token = token

	return &resp, nil
}

func Details(ctx context.Context, arg *AllRecordStruct) {}

func VerifyToken(token string) (users.TokenDetails, error) {
	details, err := SwingRepository.FetchToken(token)
	if err != nil {
		log.Println("ERROR", err)
		return users.TokenDetails{}, err
	}
	return details, nil
}

func GenerateJwt(duration time.Duration, username string, userid int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(duration * time.Minute).Unix()
	claims["authorized"] = true
	claims["user"] = username
	claims["userid"] = userid

	tokenString, err := token.SignedString(JwtSecretKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}
