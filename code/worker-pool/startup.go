package workerpool

func StartProcessing() {
	go produceWork()
	scheduler()
}
