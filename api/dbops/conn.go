package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

// Initilize tht database connection
func init() {
	dbConn, err = sql.Open("mysql", "root:19990429#@tcp(localhost:3306)/mostream?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
}

// create databases/

// example
//CREATE TABLE `users`(
//`id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
//`login_name` VARCHAR(64) NOT NULL UNIQUE KEY,
//`pwd` TEXT DEFAULT NULL
//)ENGINE InnoDB

// we do not need to create the database by go-sql-driver
