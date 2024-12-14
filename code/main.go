package main

import "main/idempotent"

// uncomment to a specific system
func main() {
	// go loadbalancer.StartLoadBalancer()
	// api.API()
	//plugin.StatisticsPlugin()
	//scaling.RunAgent()
	idempotent.Idempotent()

}
