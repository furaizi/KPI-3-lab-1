package main

import (
	"encoding/json"
	"net/http"
	"time"
	"log"
)

type TimeResponse struct {
	Time string `json:"time"`
}

func timeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	response := TimeResponse{Time: time.Now().Format(time.RFC3339)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/time", timeHandler)
	addr := ":8795"
	log.Printf("Server is running on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
