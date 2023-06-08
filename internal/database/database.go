package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ZiganshinDev/My-Pet-Projects/testForOzon/internal/models"
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "postgres"
	dbname = "OZON"
)

func createConnection() *sql.DB {
	password := os.Getenv("DB_PASSWORD")

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected")

	return db
}

func InsertLongURL(originalURL string) int64 {
	db := createConnection()
	defer db.Close()

	sqlStatement := `INSERT INTO urls (original_url) VALUES ($1) RETURNING urlid`

	var id int64

	err := db.QueryRow(sqlStatement, originalURL).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	log.Printf("Inserted a single record %v", id)

	return id
}

func GetLongURL(id int64) (string, error) {
	db := createConnection()
	defer db.Close()

	var url models.URLs

	sqlStatement := `SELECT * FROM urls WHERE urlid = $1`

	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&url.ID, &url.OriginalURL)

	if err == sql.ErrNoRows {
		log.Println("No rows were returned!")
		return url.OriginalURL, nil
	} else if err == nil {
		return url.OriginalURL, nil
	} else {
		log.Fatalf("Unable to scan the row. %v", err)
	}

	return url.OriginalURL, err
}
