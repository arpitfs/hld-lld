package workerpool

import (
	"fmt"
	"strconv"
	"time"
)

const (
	QueueSize = 5
)

type Worker struct {
	jobQueue    chan []Job
	workerCount int
}

func Start() {
	workerPool := &Worker{
		jobQueue:    make(chan []Job, QueueSize),
		workerCount: 3,
	}

	workerPool.startWorking()
}

func (w *Worker) startWorking() {
	for {
		if len(w.jobQueue) > 0 {
			job := w.jobQueue[0]
		} else {
			fmt.Println("No job to perform")
		}
		time.Sleep(5 * time.Second)
	}
}

func (w *Worker) submitJob(job Job) {
	w.jobQueue <- job
}

func (w *Worker) createJobs() {
	count := 0
	for {
		job := Job{
			id:  count,
			job: "Current Job: " + strconv.Itoa(count),
		}
		w.submitJob(job)
		count++
		time.Sleep(3 * time.Second)
	}
}
