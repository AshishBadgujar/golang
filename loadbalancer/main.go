package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(w http.ResponseWriter, r *http.Request)
}
type simpleServer struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	handleErr(err)

	return &simpleServer{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

func newLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
}

func (s *simpleServer) Address() string { return s.addr }

func (s *simpleServer) IsAlive() bool {
	// Set a timeout for the HTTP request
	client := http.Client{
		Timeout: time.Second * 5,
	}

	// Send a GET request to the server
	resp, err := client.Get(s.addr)
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	// Check if the response status code indicates success (2xx range)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		fmt.Println("Server is alive!")
		return true
	}

	fmt.Println("Server is not alive. Status code:", resp.StatusCode)
	return false
}

func (s *simpleServer) Serve(w http.ResponseWriter, r *http.Request) {
	s.proxy.ServeHTTP(w, r)
}

func (lb *LoadBalancer) getNextAvailableServer() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)] //
	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}
func (lb *LoadBalancer) serveProxy(w http.ResponseWriter, r *http.Request) {
	targetServer := lb.getNextAvailableServer()
	fmt.Printf("forwarding request to address %q\n", targetServer.Address())
	targetServer.Serve(w, r)
}

func main() {
	servers := []Server{
		newSimpleServer("https://www.facebook.com"),
		newSimpleServer("http://www.bing.com"),
		newSimpleServer("http://www.duckduckgo.com"),
	}
	lb := newLoadBalancer("8000", servers)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		lb.serveProxy(w, r)
	})

	fmt.Printf("serving requests at localhost: %s\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
