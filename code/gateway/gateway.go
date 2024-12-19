package gateway

import (
	"fmt"
	"net/http"
)

const (
	Backend1            = "http://localhost:8081"
	Backend2            = "http://localhost:8082"
	Port                = ":8080"
	HttpMethodGet       = "GET"
	BucketCapacity      = 5
	BucketRefillingRate = 10
)

func StartHandlingRequest() {
	server := requestHandlers()

	fmt.Println("API Gateway Started at port 8080")
	http.ListenAndServe(Port, server)
}
