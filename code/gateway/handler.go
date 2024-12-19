package gateway

import (
	"net/http"

	"github.com/gorilla/mux"
)

func requestHandlers() http.Handler {
	router := mux.NewRouter()
	server := setUpRateLimiter(router)
	router.HandleFunc("/github", Proxy(Backend1)).Methods(HttpMethodGet)
	router.HandleFunc("/website", Proxy(Backend2)).Methods(HttpMethodGet)

	return server
}
