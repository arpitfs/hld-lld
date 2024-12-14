package streaming

import (
	"fmt"
	"net/http"
)

func streamingHandler(w http.ResponseWriter, r *http.Request) {
	if !isProcessCompleted {
		streamLogs(w, r)
	} else {
		http.ServeFile(w, r, fileName)
	}
}

func streamLogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	for logMsg := range stramingChannel {
		fmt.Fprintf(w, "data: %s\n\n", logMsg)
		w.(http.Flusher).Flush()
	}

	w.(http.Flusher).Flush()
}
