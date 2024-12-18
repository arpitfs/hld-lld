package workerpool

import (
	"fmt"
	"sync"
	"time"
)

func manager(jobQueue chan string, wg sync.WaitGroup, workers chan struct{}) {
	for {
		time.Sleep(7 * time.Second)
		select {
		case data := <-jobQueue:
			fmt.Println("Manager picked:", data)
			select {
			case workers <- struct{}{}:
				wg.Add(1)
				go func(text string) {
					defer func() {
						<-workers
						wg.Done()
					}()
					processText(data)
				}(data)
			default:
				fmt.Println("No resources to process text:", data)
			}
		default:
			fmt.Println("No job to process!")
		}
	}
}

func processText(data string) {
	fmt.Println("Processing:", data)
	time.Sleep(3 * time.Second)
	fmt.Println("Completed processing:", data)
}
