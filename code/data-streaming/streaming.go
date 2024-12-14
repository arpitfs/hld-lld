package streaming

func Streaming() {
	go generateLogs()
	handler()
}
