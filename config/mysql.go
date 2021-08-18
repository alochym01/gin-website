package config

import (
	"database/sql"
	"fmt"
)

func MysqlConn(host string, port int64, username, password, dbname string) *sql.DB {
	dbSource := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8",
		username,
		password,
		host,
		port,
		dbname,
	)

	dbConn, err := sql.Open("mysql", dbSource)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	return dbConn
}
