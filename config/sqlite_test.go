package config

import (
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestSqliteConn(t *testing.T) {
	db := SqliteConn("../foo.db")
	err := db.Ping()
	if err != nil {
		t.Errorf("Database fail to connect with err: %s", err.Error())
	}
}
