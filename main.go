package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var (
	port = ":9090"
	c    = int32(0)
)

func main() {

	fmt.Println("hello docker")
	f := func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&c, 1)
		fmt.Fprintf(w, "welcome %d guest", c)
	}
	http.HandleFunc("/", f)
	log.Printf("listen:%s", port)
	http.ListenAndServe(port, nil)
}
