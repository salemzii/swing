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
		timestamp TIMESTAMP(6) NOT NULL,
		created TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`

	insert = `INSERT INTO records (
		message,
		level,
		stacktrace,
		function, 
		linenumber,
		offset,
		timestamp
	)`
)

/*
id BIGINT AUTO_INCREMENT PRIMARY KEY,
content varchar(300) NOT NULL,
createdate TIMESTAMP(6) NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO messages (
content
) VALUES (
'starting docker container'
);
*/
