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

type PostgreSQLStorage struct {
	db *sql.DB
}

func NewPostgreSQLStorage() *PostgreSQLStorage {
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

	return &PostgreSQLStorage{
		db: db,
	}
}

func (s *PostgreSQLStorage) InsertURLs(originalURL, shortURL string) {
	defer s.db.Close()

	sqlStatement := `INSERT INTO urls (original_url, short_url) VALUES ($1, $2)`

	err := s.db.QueryRow(sqlStatement, originalURL, shortURL)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	log.Print("Inserted a single record")
}

func (s *PostgreSQLStorage) GetLongURL(shortURL string) (string, error) {
	defer s.db.Close()

	var url models.URLs

	sqlStatement := `SELECT * FROM urls WHERE short_url = $1`

	row := s.db.QueryRow(sqlStatement, shortURL)

	err := row.Scan(&url.ID, &url.OriginalURL, &url.ShortURL)

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

func (s *PostgreSQLStorage) IsShortURLExists(shortURL string) bool {
	defer s.db.Close()

	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM urls WHERE short_url = $1)", shortURL).Scan(&exists)
	if err != nil {
		return false
	}
	return exists
}

func (s *PostgreSQLStorage) IsOriginalURLExists(originalURL string) bool {
	defer s.db.Close()

	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM urls WHERE original_url = $1)", originalURL).Scan(&exists)
	if err != nil {
		return false
	}

	return exists
}
