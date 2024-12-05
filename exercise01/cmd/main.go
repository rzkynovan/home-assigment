package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	fmt.Printf("server is starting and listening at :8080\n")
	http.ListenAndServe(":8080", nil)
}
