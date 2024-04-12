package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/stdlib"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
}

var dbConn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifetime = 5 * time.Minute

func ConnectPostgres(dsn string) (*DB, error) {

	database, err := sql.Open("pgx",dsn)

	if err != nil {
		return nil, err
	}

	database.SetMaxOpenConns(maxOpenDbConn)
	database.SetMaxIdleConns(maxIdleDbConn)
	database.SetConnMaxLifetime(maxDbLifetime)

	err = testDb(database)

	if err != nil {
		return nil, err
	}

	dbConn.DB = database
	return dbConn, nil
}


func testDb(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		fmt.Println("Error")
		return err
	}
	fmt.Println("*** ping database successfully")
	return nil
} 