package driver

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// DB holds the database connection pool
type DB struct {
	SQL *sql.DB
}

var dbconn = &DB{}

const maxOpenDbConn = 10
const maxIdleDbConn = 5
const maxDbLifeTime = 5 * time.Minute

// ConnectSQL Creates Database pool for PostgreSQL
func ConnectSQL(dsn string) (*DB, error) {
	d, err := NewDataBase(dsn)
	if err != nil {
		panic(err)
	}

	d.SetMaxOpenConns(maxOpenDbConn)
	d.SetMaxIdleConns(maxIdleDbConn)
	d.SetConnMaxLifetime(maxDbLifeTime)

	dbconn.SQL = d

	err = testDB(d)
	if err != nil {
		return nil, err
	}

	return dbconn, nil
}

// testDB Tries to ping the DataBase
func testDB(d *sql.DB) error {
	err := d.Ping()
	if err != nil {
		log.Println("Cannot ping a database")
		return err
	}

	return nil
}

// NewDataBase Creates a new database for the applictaion
func NewDataBase(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Println("cannot create a database")
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Println("Cannot ping a  new database")
		return nil, err
	}

	return db, nil
}
