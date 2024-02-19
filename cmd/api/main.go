package main

import (
	"fmt"
	"log"
	"net/http"
	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/config"
	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/database"
	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/metrics"
	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/router"
	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/util"
)

func main() {
	// Load logo
	util.Figure()

	// Load configuration
	config.Load()
	log.Println("[TRACE] Database connection string", config.StringConnectionDB)
	log.Println("[INFO] API running on port", config.AppApiPort)

	// Setup database
	log.Println("[INFO] Setting up database...")
	database.Init()

	// Load metrics setup
	log.Println("[INFO] Setting up Server metrics...")
	metrics.InitMetrics()

	// Run
	log.Println("[INFO] Running API!")
	r := router.Generate()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.AppApiPort), r))
}
