package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(5)
	// start proxy server (load balancer)
	go func() {
		Serve()
		wg.Done()
	}()

	// starting actual servers
	go func() {
		startServer(":8081")
		wg.Done()
	}()
	go func() {
		startServer(":8082")
		wg.Done()
	}()
	go func() {
		startServer(":8083")
		wg.Done()
	}()
	go func() {
		startServer(":8084")
		wg.Done()
	}()
	wg.Wait()
}

func startServer(port string) {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Server strted at port: %s\n", port)
			fmt.Fprintf(w, "Response Header: %v\n", r.Header)
		}))
	fmt.Println("Server starting at port", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		fmt.Println("Error while starting server at port:", port)
		log.Fatal(err)
	}

}
