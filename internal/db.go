package internal

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbName     = "coworking"
	dbUser     = "root"
	dbPassword = "password"
	dbPort     = "3306"
	dbHost     = "127.0.0.1"
)

func DB() *sql.DB {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Fatalf("Database connection failed, %v", err)
	}

	return db
}
