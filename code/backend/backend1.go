package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func server1() {
	http.HandleFunc("/github", backend1)

	fmt.Println("Backend Server running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func backend1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Response From Backend 1")
}
