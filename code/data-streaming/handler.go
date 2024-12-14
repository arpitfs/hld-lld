package streaming

import (
	"log"
	"net/http"
)

func handler() {
	http.HandleFunc("/streaming", streamingHandler)
	log.Println("Streaming Server started at :8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}

func streamingHandler(w http.ResponseWriter, r *http.Request) {
	if !isProcessCompleted {
		streamLogs(w, r)
	} else {
		http.ServeFile(w, r, FileName)
	}
}
