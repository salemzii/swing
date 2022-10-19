package db

const (
	migrate = `CREATE TABLE IF NOT EXISTS records (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		message varchar(300) NOT NULL, 
		level varchar(50) NOT NULL,
		stacktrace varchar(200) NOT NULL,
		function varchar(100) NOT NULL,
		linenumber INT NOT NULL,
		offset INT NOT NULL,
		timestamp TIMESTAMP(6),
		created TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	create = `INSERT INTO records (
		message,
		level,
		stacktrace,
		function, 
		linenumber,
		offset,
		timestamp
	)	VALUES(?, ?, ?, ?, ?, ?, ?);`

	insertMany       = "INSERT INTO records (message, level, stacktrace, function, linenumber, offset, timestamp)	VALUES "
	getLast15Minutes = "SELECT * FROM records WHERE created  > DATE_SUB(NOW(), INTERVAL 15 MINUTE);"

	getLastXMinutes = `SELECT * FROM records WHERE created  > DATE_SUB(NOW(), INTERVAL ? MINUTE);`

	all           = "SELECT * FROM records ORDER BY created;"
	getByFunction = `SELECT * 
	FROM records 
	WHERE function=? 
	ORDER BY created; 
	`
	getByLevel = `SELECT * 
	FROM records
	WHERE level=?
	ORDER BY created;
	`
	getByLineNum = `SELECT * FROM records WHERE linenumber=? ORDER BY created;`
)
