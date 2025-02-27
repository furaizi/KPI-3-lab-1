package main

import (
	"encoding/json" // JSON encoding
	"net/http"      // HTTP server
	"time"          // Time functions
	"log"           // Logging
)

// TimeResponse represents the JSON response with the current time
// The `json` tag ensures correct JSON field naming.
type TimeResponse struct {
	Time string `json:"time"`
}

// timeHandler handles requests to /time and returns the current time in JSON format.
// Only GET requests are allowed; others receive a "405 Method Not Allowed" response.
func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Generate the current time response in RFC3339 format.
	response := TimeResponse{Time: time.Now().Format(time.RFC3339)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {

	// Register the timeHandler function to handle /time requests.
	http.HandleFunc("/time", timeHandler)

	addr := ":8795"

	// Log server startup and listen for incoming requests.
	log.Printf("Server is running on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
