package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"sync"
	"runtime"
)

var visitorCount int
var mu sync.Mutex

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	visitorCount++
	count := visitorCount
	mu.Unlock()

	hostname, _ := os.Hostname()
	clientIP, _, _ := net.SplitHostPort(r.RemoteAddr)

	fmt.Fprintf(w, "<h1>Hello from Go 🚀</h1>")
	fmt.Fprintf(w, "<p>Server Hostname: %s</p>", hostname)
	fmt.Fprintf(w, "<p>Client IP: %s</p>", clientIP)
	fmt.Fprintf(w, "<p>Visitor number: %d</p>", count)
	fmt.Fprintf(w, "<p>Go Version: %s</p>", runtime.Version())
	fmt.Fprintf(w, "<p>OS/Arch: %s/%s</p>", runtime.GOOS, runtime.GOARCH)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started at http://localhost:8099")
	if err := http.ListenAndServe(":8099", nil); err != nil {
		panic(err)
	}
}
