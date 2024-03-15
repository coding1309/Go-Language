package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type logHandler struct {
	handler http.Handler
}

func (l *logHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	duration := time.Since(start)
	log.Printf("%s %s %v", r.RemoteAddr, r.URL, duration)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	http.Handle("/", &logHandler{handler: http.HandlerFunc(helloHandler)})
	http.HandleFunc("/private", privateHandler)

	fmt.Println("Server listening on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func privateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Private: Welcome, %s!", r.URL.Path[1:])
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// Check it Out

// http://localhost:8080/
// http://localhost:8080/neel
// http://localhost:8080/private/patel
