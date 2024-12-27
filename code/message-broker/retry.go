package messagebroker

import (
	"fmt"
	"time"
)

const (
	MaxRetry = 2
)

type Retry struct {
	retryCount int
	message    string
}

func retryMessage(retry Retry) {
	if retry.retryCount < MaxRetry {
		fmt.Printf("Retrying message %s and count %d : ", retry.message, retry.retryCount)
		fmt.Println()
		retry.retryCount++
		time.Sleep(time.Second * time.Duration(retry.retryCount))
		retryMessage(retry)
	} else {
		fmt.Println("Message moved to dead letter queue: ", retry.message)
		deadLetterQueue <- retry
	}
}
