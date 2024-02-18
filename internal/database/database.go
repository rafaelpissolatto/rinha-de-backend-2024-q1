package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"rinha-backend-2024q1/internal/domain/entities"
	"time"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/mattn/go-sqlite3"
)

var (
	dburl = os.Getenv("DB_URL")
)

// Init initializes the database and creates the dabases and tables if they do not exist
func Init() {
	migrate()
}

// createDatabase creates the database if it does not exist
func migrate() {
	db, err := sql.Open("sqlite3", os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()

	// CREATE TABLE IF NOT EXISTS customer based on the customerModel
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS customers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		limit INTEGER,
		balance INTEGER
	)`)
	if err != nil {
		log.Fatal(err.Error())
	}

	// CREATE TABLE IF NOT EXISTS transaction based on the transactionModel
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS transaction (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		customer_id INTEGER,
		value INTEGER,
		type TEXT,
		description TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		log.Fatal(err.Error())
	}

	// Seed the database
	_, err = db.Exec(`INSERT INTO transactions (client_id, amount, created_at) 
		VALUES (1, 1000, datetime('now'))`)
	if err != nil {
		log.Fatal(err.Error())
	}

}

type Service interface {
	Health() map[string]string
}

type service struct {
	db *sql.DB
}

func New() Service {
	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		// This will not be a connection error, but a DSN parse error or
		// another initialization error.
		log.Fatal(err)
	}
	s := &service{db: db}
	return s
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.db.PingContext(ctx)
	if err != nil {
		log.Fatalf(fmt.Sprintf("db down: %v", err))
	}

	return map[string]string{
		"message": "It's healthy",
	}
}

func (s *service) Close() {
	s.db.Close()
}

func GetExtract(id int) (entities.Transaction, error) {
	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var t entities.Transaction
	err = db.QueryRow("SELECT * FROM transactions WHERE id = ?", id).Scan(&t.ID, &t.CustomerID, &t.Value, &t.Type, &t.Description, &t.CreatedAt)
	if err != nil {
		return t, err
	}
	return t, nil
}

func GetExtracts() (entities.Transactions, error) {
	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM transactions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ts entities.Transactions
	for rows.Next() {
		var t entities.Transaction
		err = rows.Scan(&t.ID, &t.CustomerID, &t.Value, &t.Type, &t.Description, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		ts = append(ts, t)
	}
	return ts, nil
}

func CreateExtract(t entities.Transaction) (entities.Transaction, error) {
	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO transactions (client_id, amount, created_at) VALUES (?, ?, ?)", t.CustomerID, t.Value, t.CreatedAt)
	if err != nil {
		return t, err
	}
	return t, nil
}

func UpdateExtract(id int, t entities.Transaction) (entities.Transaction, error) {
	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE transactions SET client_id = ?, amount = ?, created_at = ? WHERE id = ?", t.CustomerID, t.Value, t.CreatedAt, id)
	if err != nil {
		return t, err
	}
	return t, nil
}

func 