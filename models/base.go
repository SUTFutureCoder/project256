package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"project256/util"
	"fmt"
)

var mysqlConn *sql.DB

func GetDbConn() *sql.DB {
	if mysqlConn == nil {
		cfg, err := util.GetConfig("mysql")
		mysqlConn, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg["username"], cfg["password"], cfg["host"], cfg["port"], cfg["default_db"]))
		if err != nil {
			log.Fatal(err)
			return nil
		}
	}
	return mysqlConn
}