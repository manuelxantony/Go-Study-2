package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

var cfg Config
var mu sync.Mutex
var idx int = 0

func lbHandler(w http.ResponseWriter, r *http.Request) {
	maxLen := len(cfg.Backend)
	// round robin
	mu.Lock()
	urlPort := cfg.Backend[idx%maxLen].URL
	targetURL, err := url.Parse(urlPort)
	if err != nil {
		log.Fatal(err)
	}
	idx++
	mu.Unlock()
	// redirecting to target url
	resverseProxy := httputil.NewSingleHostReverseProxy(targetURL)
	resverseProxy.ServeHTTP(w, r)
}

func Serve() {
	// reverse proxy
	// director

	cfg = readConfig()

	s := &http.Server{
		Addr:    ":" + cfg.Proxy.Port,
		Handler: http.HandlerFunc(lbHandler),
	}
	fmt.Println("Proxy Server started at", s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal("Error stasting proxy server at port:", s.Addr, err)
	}

}
