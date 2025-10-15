package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World from Go!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started at http://localhost:8099")
	if err := http.ListenAndServe(":8099", nil); err != nil {
		panic(err)
	}
}
