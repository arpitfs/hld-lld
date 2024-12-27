package messagebroker

import (
	"fmt"
	"time"
)

func processMessages() {
	for msg := range messageQueue {
		if msg == "dl1" {
			retyMessage := Retry{
				retryCount: 0,
				message:    msg,
			}

			go retryMessage(retyMessage)

		} else {
			fmt.Println("Processing message: ", msg)
			time.Sleep(5 * time.Second)
		}
	}
}
