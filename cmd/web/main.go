package main

import (
	"os"

	"github.com/alochym01/gin-website/config"
	"github.com/alochym01/gin-website/models"
	"github.com/alochym01/gin-website/router"

	// _ "github.com/go-sql-driver/mysql" // 172.17.0.2
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Remove sqlite db file
	os.Remove("foo.db")
	config.DB = config.SqliteConn("foo.db")
	defer config.DB.Close()
	config.DB.Ping()
	models.PreparesqliteDB(config.DB)

	// // Mysql
	// config.DB = config.MysqlConn("172.17.0.2", 3306, "cmis", "Phuongtt@123cmis", "MYSQLTEST")
	// defer config.DB.Close()
	// config.DB.Ping()
	// models.PrepareMysqlDB(config.DB)

	router := router.Router()

	router.Run()
}
