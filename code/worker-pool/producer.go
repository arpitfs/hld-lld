package workerpool

import (
	"fmt"
	"sync"
	"time"
)

var producer Producer
var work = 1

type Producer struct {
	mu   sync.Mutex
	work []int
	cap  int
}

func produceWork() {
	time.Sleep(2 * time.Second)
	producer = Producer{cap: 4}
	for {
		producer.mu.Lock()
		if len(producer.work) < 4 {
			producer.work = append(producer.work, work)
			fmt.Printf("New work no %d created\n", work)
			work++
		} else {
			fmt.Println("Queue is full for work no: ", work)
		}
		producer.mu.Unlock()
		time.Sleep(time.Second * 1)
	}
}
