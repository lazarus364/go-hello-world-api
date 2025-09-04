package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// response is a simple struct we will encode to JSON
type response struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timestamp"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Set JSON content type
	w.Header().Set("Content-Type", "application/json")

	// Create the response payload
	payload := response{
		Message:   "Hello, World!",
		Timestamp: time.Now(),
	}

	// Encode to JSON and write to the response
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	// Map the root path to our handler
	http.HandleFunc("/", helloHandler)

	// Start the server on port 8080
	log.Println("Server starting on http://localhost:8080 …")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
