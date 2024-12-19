package backend

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func server2() {
	http.HandleFunc("/website", backend2)

	fmt.Println("Backend Server running on http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", nil))
}

func backend2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Response From Backend 2")
}
