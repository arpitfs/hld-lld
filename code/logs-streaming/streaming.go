package streaming

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var stramingChannel = make(chan string, 100)
var isProcessCompleted = false

const (
	fileName = "streaming.log"
)

func Streaming() {
	streamingLogFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening a file:", err)
		return
	}
	defer streamingLogFile.Close()

	go generateLogs(streamingLogFile)

	http.HandleFunc("/streaming", streamingHandler)
	log.Println("Streaming Server started at :8083")
	log.Fatal(http.ListenAndServe(":8083", nil))
}

func generateLogs(logFile *os.File) {
	logger := log.New(logFile, "", log.Ldate|log.Ltime)
	for i := 1; i <= 10; i++ {
		message := fmt.Sprintf("%d ) Processing Streaming", i)
		logger.Println(message)
		stramingChannel <- message
		time.Sleep(2 * time.Second)
	}
	close(stramingChannel)
	isProcessCompleted = true
}
