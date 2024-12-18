package workerpool

import "sync"

const (
	QueueSize  = 5
	MaxThreads = 3
)

func StartProcessing() {
	jobQueue := make(chan string, QueueSize)
	var wg sync.WaitGroup
	workers := make(chan struct{}, MaxThreads)

	go createJobs(jobQueue)
	go manager(jobQueue, wg, workers)

	select {}
}
