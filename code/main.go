package main

import parkinglot "main/parking-lot"

// uncomment to a specific system
func main() {
	// go loadbalancer.StartLoadBalancer()
	// api.API()
	//plugin.StatisticsPlugin()
	//scaling.RunAgent()
	//idempotent.Idempotent()
	//streaming.Streaming()
	// go backend.Backend()
	// gateway.StartHandlingRequest()
	parkinglot.StartParking()
}
