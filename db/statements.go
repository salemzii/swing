package db

const (
	migrate = `CREATE TABLE IF NOT EXISTS records (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		message varchar(300) NOT NULL, 
		level varchar(50) NOT NULL,
		stacktrace varchar(200) NOT NULL,
		function varchar(100) NOT NULL,
		linenumber INT NOT NULL,
		process INT NOT NULL,
		timestamp TIMESTAMP(6),
		logger varchar(100) NOT NULL,
		created TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`

	create = `INSERT INTO records (
		message,
		level,
		stacktrace,
		function, 
		linenumber,
		process,
		timestamp,
		logger
	)	VALUES(?, ?, ?, ?, ?, ?, ?, ?);`

	insertMany       = "INSERT INTO records (message, level, stacktrace, function, linenumber, process, timestamp, logger)	VALUES "
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
	delete       = `DELETE FROM records WHERE id=?;`
	deleteMany   = `DELETE FROM records WHERE id IN `

	GetToken    = "SELECT FROM tokens WHERE token = ?"
	migrateUser = `CREATE TABLE IF NOT EXISTS users (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		email VARCHAR(50)	NOT NULL,
		password VARCHAR(50) NOT NULL,
		created TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
	`
	migrateToken = `CREATE TABLE IF NOT EXISTS tokens (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		token VARCHAR(20) NOT NULL, 
		expires TIMESTAMP(6) NOT NULL,
		rate INT NOT NULL,
		enabled BOOLEAN NOT NULL,
		created TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP,
		userid BIGINT UNIQUE FOREIGN KEY REFERENCES users(id)
	);
	`
)
