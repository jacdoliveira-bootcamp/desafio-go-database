package testutils

import (
	"github.com/DATA-DOG/go-txdb"
	"github.com/go-sql-driver/mysql"
)

func RegisterDatabase() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Net:    "tcp",
		Addr:   "localhost:3306",
		DBName: "fantasy_products_test",
	}
	txdb.Register("txdb", "mysql", cfg.FormatDSN())
}
