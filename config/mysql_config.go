package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	db *sql.DB
)

func init() {
	//dataSourceName := "root:YPlove156520.@tcp(localhost:3306)/yp_nav?charset=utf8mb4"
	dataSourceName := "root:YPlove156520.@tcp(114.115.161.107:3306)/yp_nav?charset=utf8mb4"
	connection, e := sql.Open("mysql", dataSourceName)
	if e != nil {
		log.Println(e)
	}
	db = connection
}

func GetDB() *sql.DB {
	return db
}
