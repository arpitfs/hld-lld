package gateway

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func Proxy(instance string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url, err := url.Parse(instance)
		if err != nil {
			http.Error(w, "Bad Gateway", http.StatusBadGateway)
			fmt.Printf("Error parsing instance URL: %v", err)
			return
		}

		proxy := httputil.NewSingleHostReverseProxy(url)
		r.Host = url.Host
		proxy.ServeHTTP(w, r)
	}
}
