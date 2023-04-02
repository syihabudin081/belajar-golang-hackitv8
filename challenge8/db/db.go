package db

import (
    "database/sql"
    "fmt"

    _ "github.com/lib/pq"
)

const(
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "123"
	db_name = "books_db"
)

// NewDB returns a new database connection
func NewDB() (*sql.DB, error) {
    connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, db_name)
    return sql.Open("postgres", connStr)
}