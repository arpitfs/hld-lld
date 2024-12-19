package main

import (
	"main/backend"
	"main/gateway"
)

// uncomment to a specific system
func main() {
	// go loadbalancer.StartLoadBalancer()
	// api.API()
	//plugin.StatisticsPlugin()
	//scaling.RunAgent()
	//idempotent.Idempotent()
	//streaming.Streaming()
	go backend.Backend()
	gateway.StartHandlingRequest()
}
