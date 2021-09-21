package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello, world (from go!)\n"))
	})
	addr := ":80"
    fmt.Println("HTTP service listening on port:", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
