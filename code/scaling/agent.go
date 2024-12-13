package scaling

import (
	"time"
)

func RunAgent() {
	for {
		monitorServices()
		time.Sleep(MonitoringTimePeriod * time.Second)
	}
}
