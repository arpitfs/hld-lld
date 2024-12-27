package main

import messagebroker "main/message-broker"

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
	// parkinglot.StartParking()

	messagebroker.StartServiceBus()
}
