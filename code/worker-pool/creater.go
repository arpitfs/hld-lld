package workerpool

import (
	"fmt"
	"time"
)

type Job struct {
	id  int
	job string
}

func createJobs(jobQueue chan string) {
	count := 0
	for {
		data := fmt.Sprintf("Data to be processed %d", count)
		select {
		case jobQueue <- data:
			fmt.Println("Job Generated: ", count)
		default:
			fmt.Println("Queue is full")
		}
		count++
		time.Sleep(time.Second * 5)
	}
}
