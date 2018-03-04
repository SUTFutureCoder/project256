package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var mysqlConn *sql.DB

func GetDbConn() *sql.DB {
	if mysqlConn == nil {
		mysqlConn, err := sql.Open("mysql", "root:000000@tcp(127.0.0.1:3306)/project256")
		if err != nil {
			log.Fatal(err)
		}
		defer mysqlConn.Close()
	}
	return mysqlConn
}