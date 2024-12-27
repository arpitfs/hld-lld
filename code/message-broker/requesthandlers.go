package messagebroker

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler() {
	http.HandleFunc("/broker", messageHandler)
	http.HandleFunc("/dlcount", dlHandler)
	log.Println("Message Broker Started at :8085")
	log.Fatal(http.ListenAndServe(":8085", nil))
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	messageKey := r.Header.Get("message_key")
	fmt.Println("Message Received: ", messageKey)

	messageQueue <- messageKey
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Message created successfully")

}

func dlHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	dlCount := len(deadLetterQueue)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dlCount)

}
