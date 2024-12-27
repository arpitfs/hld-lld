package messagebroker

var messageQueue = make(chan string, 10)
var deadLetterQueue = make(chan Retry, 10)
