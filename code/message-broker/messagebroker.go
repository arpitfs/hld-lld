package messagebroker

func StartServiceBus() {
	go processMessages()
	handler()
}
