package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// stringConnectionDB is the string connection to the database
	StringConnectionDB = ""

	// APIPort is the port that the API will run
	AppApiPort = 0
)

// Load loads the configuration from the environment variables
func Load() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal("[ERROR] Failed to load .env file", err)
	}

	AppApiPort, err = strconv.Atoi(os.Getenv("APP_API_PORT"))
	if err != nil {
		log.Println("[ERROR] Failed to load APP_API_PORT from .env file, using default value 8081")
		AppApiPort = 8081
	}

	// string connection for sqlite
	StringConnectionDB = fmt.Sprintf("%s", os.Getenv("APP_DB_PATH")) // "rinha.db"
}
