package gateway

import (
	"fmt"
	"net/http"
)

func StartHandlingRequest() {
	server := requestHandlers()

	fmt.Println("API Gateway Started at port 8080")
	http.ListenAndServe(Port, server)
}
