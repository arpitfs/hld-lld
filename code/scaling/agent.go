package scaling

import (
	"time"
)

func RunAgent() {
	for {
		monitorServices()
		time.Sleep(10 * time.Second)
	}
}
