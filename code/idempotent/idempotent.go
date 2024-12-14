package idempotent

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var IdempotentTracker = make(map[string]bool)

func idempotent() {
	http.HandleFunc("/idempptent", getPaginatedData)
	fmt.Println("Idempotent Server running on http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func getPaginatedData(w http.ResponseWriter, r *http.Request) {
	if r.
	idempotentKey := r.Header.Get("idempotent_key")

	if _, exists := IdempotentTracker[idempotentKey]; exists {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Is not present")
	}

}
