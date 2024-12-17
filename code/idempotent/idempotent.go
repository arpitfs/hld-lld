package idempotent

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var IdempotentTracker = make(map[string]bool)

func Idempotent() {
	http.HandleFunc("/idempotent", processRequest)
	fmt.Println("Idempotent Server running on http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func processRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	idempotentKey := r.Header.Get("idempotent_key")
	fmt.Println(idempotentKey, IdempotentTracker)
	if _, exists := IdempotentTracker[idempotentKey]; !exists {
		IdempotentTracker[idempotentKey] = true
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Request Processed Successfully")
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusTooManyRequests)
		json.NewEncoder(w).Encode("Duplicate Request")
	}

}
