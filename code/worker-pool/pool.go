package workerpool

import (
	"fmt"
	"sync"
	"time"
)

const (
	MaxThreads = 2
)

func scheduler() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, MaxThreads)
	for {
		producer.mu.Lock()
		if len(producer.work) < producer.cap {
			if len(ch) < MaxThreads {
				ch <- struct{}{}
				wg.Add(1)
				go worker(&wg, ch)
			} else {
				fmt.Println("No resources available")
			}
		} else {
			fmt.Println("No work to do")
		}
		producer.mu.Unlock()
		time.Sleep(1 * time.Second)
	}
}
