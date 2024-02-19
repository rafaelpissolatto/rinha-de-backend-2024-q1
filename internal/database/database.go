package database

import (
	"database/sql"
	"log"

	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/config"

	_ "modernc.org/sqlite"
)

// Init initializes the database and creates the dabases and tables if they do not exist
func Init() {
	createTableCustomers()
	createTableTransactions()
	loadInitialData()
}

func createTableCustomers() {
	db, err := Connect()
	if err != nil {
		log.Println("[ERROR] Failed to connect to the database", err)
		return
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS customers (
		id INTEGER PRIMARY KEY,
		credit_limit INT NOT NULL,
		balance INT NOT NULL
	)`)
	if err != nil {
		log.Println("[ERROR] Failed to create the customers table", err)
		return
	}

	log.Println("[INFO] Customers table created")
}

func createTableTransactions() {
	db, err := Connect()
	if err != nil {
		log.Println("[ERROR] Failed to connect to the database", err)
		return
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY,
		customer_id INTEGER NOT NULL,
		amount INT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (customer_id) REFERENCES customers(id)
	)`)
	if err != nil {
		log.Println("[ERROR] Failed to create the transactions table", err)
		return
	}

	log.Println("[INFO] Transactions table created")
}

func loadInitialData() {
	db, err := Connect()
	if err != nil {
		log.Println("[ERROR] Failed to connect to the database", err)
		return
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO customers (credit_limit, balance) VALUES
		(100000, 0),
		(80000, 0),
		(1000000, 0),
		(10000000, 0),
		(500000, 0)
	`)
	if err != nil {
		log.Println("[ERROR] Failed to insert initial data into the customers table", err)
		return
	}

	log.Println("[INFO] Initial data loaded into the customers table")

}

// Connect opens a connection with the database
func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", config.StringConnectionDB)
	if err != nil {
		log.Println("[ERROR] Failed to connect to the database", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		log.Println("[ERROR] Failed to ping the database", err)
		return nil, err
	}

	return db, nil
}
