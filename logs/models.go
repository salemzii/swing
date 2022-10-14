package logs

import "time"

type LogRecord struct {
	Id         int       `json:"id"`
	Level      string    `json:"level"`
	Message    string    `json:"message"`
	StackTrace string    `json:"stacktrace"`
	Function   string    `json:"function"`
	LineNumber int       `json:"linenum"`
	Offset     int       `json:"offset"`
	TimeStamp  time.Time `json:"timestamp"`
	Created    time.Time `json:"created"`
}
