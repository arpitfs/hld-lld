package gateway

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func StartHandlingRequest() {
	router := mux.NewRouter()

	tokenBucket := NewTokenBucket(5, 10*time.Second)
	rateLimitMiddleware := rateLimitor(tokenBucket, router)
	router.HandleFunc("/github", Proxy("http://localhost:8081")).Methods("GET")
	router.HandleFunc("/website", Proxy("http://localhost:8082")).Methods("GET")

	fmt.Println("API Gateway Started at port 8080")
	http.ListenAndServe(":8080", rateLimitMiddleware)
}
