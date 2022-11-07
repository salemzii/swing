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
		tokenid varchar(50) NOT NULL,
		userid INT NOT NULL,
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
		logger,
		userid
	)	VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?);`

	insertMany       = `INSERT INTO records (message, level, stacktrace, function, linenumber, process, timestamp, logger, userid)	VALUES `
	getLast15Minutes = `SELECT 
	id, level, message, stacktrace, function,linenumber, process, timestamp, created, logger, userid
	FROM records WHERE userid=? AND created  > DATE_SUB(NOW(), INTERVAL 15 MINUTE) 
	LIMIT 500;`

	getLastXMinutes = `SELECT 
	id, level, message, stacktrace, function,linenumber, process, timestamp, created, logger, userid
	FROM records WHERE userid=? AND created  > DATE_SUB(NOW(), INTERVAL ? MINUTE) 
	LIMIT 500;`

	all = `SELECT 
	id, level, message, stacktrace, function,linenumber, process, timestamp, created, logger, userid
	FROM records WHERE userid=? 
	LIMIT 500;`

	getByFunction = `SELECT 
	id, level, message, stacktrace, function,linenumber, process, timestamp, created, logger, userid
	FROM records 
	WHERE userid=? AND function=? 
	LIMIT 500; 
	`
	getByLevel = `SELECT 
	id, level, message, stacktrace, function,linenumber, process, timestamp, created, logger, userid
	FROM records
	WHERE userid=? AND level=?
	LIMIT 500;
	`
	getByLineNum = `SELECT 
	id, level, message, stacktrace, function,linenumber, process, timestamp, created, logger, userid
	FROM records 
	WHERE userid=? AND linenumber=? 
	LIMIT 500;
	`
	delete     = `DELETE FROM records WHERE userid=? AND id=?;`
	deleteMany = `DELETE FROM records WHERE userid=? AND id IN `

	migrateUser = `CREATE TABLE IF NOT EXISTS users (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		email VARCHAR(50)	NOT NULL,
		password VARCHAR(50) NOT NULL,
		created TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`
	createUser = `INSERT INTO users (
		username, 
		email, 
		password
	) VALUES (?, ?, ?);`

	allUsers          = `SELECT * FROM users ORDER BY created;`
	getUserByEmail    = `SELECT id, email, password, username FROM users WHERE email=?;`
	getUserByUsername = `SELECT * FROM users WHERE username=?;`

	migrateToken = `CREATE TABLE IF NOT EXISTS tokens (
		id BIGINT AUTO_INCREMENT PRIMARY KEY,
		token VARCHAR(20) NOT NULL, 
		expires TIMESTAMP(6) NOT NULL,
		rate INT NOT NULL,
		enabled BOOLEAN NOT NULL,
		userid INT NOT NULL,
		created TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
	);`
	createToken = `INSERT INTO tokens (
		token,
		expires,
		rate, 
		enabled, 
		userid
	) VALUES(?, ?, ?, ?, ?)`
	getToken         = `SELECT * FROM tokens WHERE token = ?;`
	getTokenByUserId = `SELECT * FROM tokens WHERE userid=? RETURNING token;`
)
