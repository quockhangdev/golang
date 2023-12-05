package main

import (
	"log"
	"net/http"
	"os"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World!"))
	if err != nil {
		return
	}
}

func main() {
	addr := os.Getenv("ADDR")

	if len(addr) == 0 {
		addr = ":8080"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", HelloHandler)

	log.Printf("Server is listening at %s...", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
