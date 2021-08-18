package config

import "database/sql"

func SqliteConn(f string) *sql.DB {
	dbConn, err := sql.Open("sqlite3", "./foo.db")

	if err != nil {
		panic(err)
	}

	return dbConn
}
