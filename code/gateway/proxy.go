package gateway

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Proxy(instance string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the target URL
		url, err := url.Parse(instance)
		if err != nil {
			http.Error(w, "Bad Gateway", http.StatusBadGateway)
			fmt.Printf("Error parsing target URL: %v", err)
			return
		}

		// Create a reverse proxy
		proxy := httputil.NewSingleHostReverseProxy(url)

		// Update the request's host header to match the target server
		r.Host = url.Host

		// Forward the request to the target
		proxy.ServeHTTP(w, r)
	}
}
