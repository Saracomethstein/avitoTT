package openapi

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

func SetupDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DATABASE")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	sqldb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	err = sqldb.Ping()
	if err != nil {
		log.Fatalf("Failed to ping the database: %v", err)
	}

	log.Println("Successfully connected to the database")
	return sqldb
}
