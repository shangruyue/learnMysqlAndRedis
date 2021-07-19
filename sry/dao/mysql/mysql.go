package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/stdlib"
)

var db *sql.DB

func InitMysql()(db *sql.DB) {
	//mysqlStr := "postgres://postgres:AWc7nCRAeesaa3d7C6NS@uservicedb-dev.cyccjyevtjgf.us-west-2.rds.amazonaws.com:5432/uservicedbdev"

	mysqlStr := "root:root@tcp(127.0.0.1:3306)/mysql"
	//db, err := sql.Open("pgx", mysqlStr)
	db, err := sql.Open("mysql", mysqlStr)
	if err != nil {
		fmt.Printf("Open db failed: %v\n", err)
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Ping failed: %v\n", err)
	}
	return db
}