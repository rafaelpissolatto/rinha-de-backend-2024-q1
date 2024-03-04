package database

import (
	"database/sql"
	"log"

	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/config"
	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/model"

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
		log.Println("[ERROR] Failed to connect to the database: [", err, "]")
		return
	}
	defer db.Close()

	// ID start from 1 and is autoincrement
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS customers (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		credit_limit INT NOT NULL DEFAULT 0,
		balance INT NOT NULL DEFAULT 0
	)`)
	if err != nil {
		log.Println("[ERROR] Failed to create the customers table: [", err, "]")
		return
	}

	log.Println("[INFO] Customers table created")
}

func createTableTransactions() {
	db, err := Connect()
	if err != nil {
		log.Println("[ERROR] Failed to connect to the database: [", err, "]")
		return
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		customer_id INTEGER NOT NULL,
		amount INT NOT NULL,
		operator_type TEXT NOT NULL,
		description TEXT,
		created_at TEXT NOT NULL
	)`)
	if err != nil {
		log.Println("[ERROR] Failed to create the transactions table: [", err, "]")
		return
	}

	log.Println("[INFO] Transactions table created")
}

func loadInitialData() {
	db, err := Connect()
	if err != nil {
		log.Println("[ERROR] Failed to connect to the database: [", err, "]")
		return
	}
	defer db.Close()

	initialData := []model.Customer{
		{ID: 1, CreditLimit: 100000, Balance: 0},
		{ID: 2, CreditLimit: 80000, Balance: 0},
		{ID: 3, CreditLimit: 1000000, Balance: 0},
		{ID: 4, CreditLimit: 10000000, Balance: 0},
		{ID: 5, CreditLimit: 500000, Balance: 0},
	}

	tx, err := db.Begin()
	if err != nil {
		log.Println("[ERROR] Failed to start a transaction: [", err, "]")
		return
	}

	existsStmt, err := tx.Prepare("SELECT id FROM customers WHERE id = ?")
	if err != nil {
		log.Println("[ERROR] Failed to prepare the statement: [", err, "]")
		return
	}

	for _, customer := range initialData {
		var id int
		err = existsStmt.QueryRow(customer.ID).Scan(&id)
		if err == nil {
			continue
		}

		_, err = tx.Exec("INSERT INTO customers (id, credit_limit, balance) VALUES (?, ?, ?)", customer.ID, customer.CreditLimit, customer.Balance)
		if err != nil {
			log.Println("[ERROR] Failed to execute the statement: [", err, "]")
			return
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println("[ERROR] Failed to commit the transaction: [", err, "]")
		return
	}

	log.Println("[INFO] Initial data loaded into the customers table")
}

// Connect opens a connection with the database
func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite", config.StringConnectionDB)
	if err != nil {
		log.Println("[ERROR] Failed to connect to the database: [", err, "]")
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		log.Println("[ERROR] Failed to ping the database: [", err, "]")
		return nil, err
	}

	return db, nil
}
