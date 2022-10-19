package analytics

const (
	XLogLevelsDuration         = `SELECT * FROM records WHERE level=? AND created  > DATE_SUB(NOW(), INTERVAL ? MINUTE) ORDER BY created;`
	MostLoggedLevel            = `SELECT level, COUNT(level) AS ? FROM records ORDER BY created;`
	AverageErrorTime           = ``
	LineNumberLogLevel         = `SELECT * FROM records WHERE linenumber=? AND level=? RETURNING linenumber, level, COUNT(level) ORDER BY created;`
	HighestLineNumberWithLevel = ``
)
