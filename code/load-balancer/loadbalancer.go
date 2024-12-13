package loadbalancer

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync/atomic"
)

type Backend struct {
	url     string
	isAlive bool
	proxy   *httputil.ReverseProxy
}

type LoadBalancer struct {
	Backends []*Backend
	current  uint32
}

func CreateBackend(server string) *Backend {
	serverUrl, _ := url.Parse(server)
	return &Backend{
		url:     server,
		isAlive: true,
		proxy:   httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func NewLoadBalancer(backends []*Backend) *LoadBalancer {
	return &LoadBalancer{
		Backends: backends,
	}
}

func (lb *LoadBalancer) getServer() *Backend {
	for i := 0; i < len(lb.Backends); i++ {
		index := atomic.AddUint32(&lb.current, 1) % uint32(len(lb.Backends))
		backend := lb.Backends[index]
		if backend.isAlive {
			return backend
		}
	}
	return nil
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lb.getServer().proxy.ServeHTTP(w, r)
}

func StartLoadBalancer() {
	backends := []*Backend{
		CreateBackend("https://github.com/arpitfs"),
		CreateBackend("https://arpitfs.github.io/portfolio"),
		CreateBackend("https://arpitfs.medium.com/"),
	}

	lb := NewLoadBalancer(backends)

	// Perfrom health check periodically

	fmt.Println("Load Balancer read to server at 8080")
	log.Fatal(http.ListenAndServe(":8080", lb))
}
