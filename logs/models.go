package logs

type LogRecord struct {
	Id         int
	Level      string
	Message    string
	StackTrace string
	Function   string
	LineNumber int
	Offset     int
}
