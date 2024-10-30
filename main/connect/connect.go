package connect

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

type DBConnection struct {
	Db *sql.DB
}

var (
	instance *DBConnection
	once     sync.Once
)

func Connect() (*DBConnection, error) {
	once.Do(func() {
		// Retrieve the connection details from environment variables
		host := "postgres-service" // The service name of the PostgreSQL instance
		port := "5432"
		user := os.Getenv("POSTGRES_USER")
		password := os.Getenv("POSTGRES_PASSWORD")
		dbname := "your_database"

		// Construct the connection string
		connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)

		// Open a connection to the database
		db, err := sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal(err)
		}
		log.Print("db connected", db)
		// Return the database connection
		instance = &DBConnection{Db: db}
	})
	return instance, nil
}
