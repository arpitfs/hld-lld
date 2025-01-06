package workerpool

import (
	"fmt"
	"sync"
	"time"
)

func worker(wg *sync.WaitGroup, channel chan struct{}) {
	defer wg.Done()
	producer.mu.Lock()
	fmt.Println("Work Done for : ", producer.work[0])
	producer.work = producer.work[1:]
	producer.cap = producer.cap - 1
	producer.mu.Unlock()
	time.Sleep(2 * time.Second)
	<-channel
}
