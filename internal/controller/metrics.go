package controller

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"rinha-backend-2024q1-rafael-pissolatto-nunes/internal/middleware"
)

// Metrics is a controller that returns the metrics of the server
func Metrics(w http.ResponseWriter, r *http.Request) {
	// Get the number of goroutines
	numGoroutines := runtime.NumGoroutine()

	// Get the uptime
	uptime := time.Since(middleware.StartTime)

	// Get the memory stats
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	// Write the response
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("# HELP rinha_backend_2024q1_rafael_pissolatto_nunes_goroutines_total Number of goroutines running\n"))
	w.Write([]byte("# TYPE rinha_backend_2024q1_rafael_pissolatto_nunes_goroutines_total gauge\n"))
	w.Write([]byte("# HELP rinha_backend_2024q1_rafael_pissolatto_nunes_uptime_seconds Uptime of the server in seconds\n"))
	w.Write([]byte("# TYPE rinha_backend_2024q1_rafael_pissolatto_nunes_uptime_seconds gauge\n"))
	w.Write([]byte("# HELP rinha_backend_2024q1_rafael_pissolatto_nunes_memory_alloc_bytes Memory allocated\n"))
	w.Write([]byte("# TYPE rinha_backend_2024q1_rafael_pissolatto_nunes_memory_alloc_bytes gauge\n"))
	w.Write([]byte("# HELP rinha_backend_2024q1_rafael_pissolatto_nunes_memory_sys_bytes Memory obtained from the system\n"))
	w.Write([]byte("# TYPE rinha_backend_2024q1_rafael_pissolatto_nunes_memory_sys_bytes gauge\n"))

	w.Write([]byte(fmt.Sprintf("rinha_backend_2024q1_rafael_pissolatto_nunes_goroutines_total %d\n", numGoroutines)))
	w.Write([]byte(fmt.Sprintf("rinha_backend_2024q1_rafael_pissolatto_nunes_uptime_seconds %f\n", uptime.Seconds())))
	w.Write([]byte(fmt.Sprintf("rinha_backend_2024q1_rafael_pissolatto_nunes_memory_alloc_bytes %d\n", mem.Alloc)))
	w.Write([]byte(fmt.Sprintf("rinha_backend_2024q1_rafael_pissolatto_nunes_memory_sys_bytes %d\n", mem.Sys)))
}
