package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	//HOST is the database host
	HOST = "database"
	//PORT is the port of the database
	PORT = 5432
)

//ErrNoMatch is the error showing no matched record found
var ErrNoMatch = fmt.Errorf("No matching record")

//Database is the client communicating to the database
type Database struct {
	Conn *sql.DB
}

//Initialize is to connect to the database
func Initialize(username, password, database string) (Database, error) {
	db := Database{}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		HOST, PORT, username, password, database)

	conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}
	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")

	return db, nil
}
