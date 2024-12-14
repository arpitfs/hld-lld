package streaming

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	FileName     = "streaming.log"
	Capacity     = 100
	ContentType  = "text/event-stream"
	CacheControl = "no-cache"
	Connection   = "keep-alive"
)

var streamingChannel = make(chan string, Capacity)
var isProcessCompleted = false

func streamLogs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ContentType)
	w.Header().Set("Cache-Control", CacheControl)
	w.Header().Set("Connection", Connection)

	for stream := range streamingChannel {
		fmt.Fprintf(w, "data: %s\n\n", stream)
		w.(http.Flusher).Flush()
	}

	w.(http.Flusher).Flush()
}

func generateLogs() {
	streamingLogFile, err := os.OpenFile(FileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening a file:", err)
		return
	}
	defer streamingLogFile.Close()
	logger := log.New(streamingLogFile, "", log.Ldate|log.Ltime)
	for i := 1; i <= Capacity; i++ {
		message := fmt.Sprintf("%d ) Streaming", i)
		logger.Println(message)
		streamingChannel <- message
		time.Sleep(2 * time.Second)
	}
	close(streamingChannel)
	isProcessCompleted = true
}
