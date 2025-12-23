package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var logfile *os.File

func main() {
	var err error
	logfile, err = os.OpenFile("/app/data/app.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644)

	if err != nil {
		log.Println("application error in log", err)
	}

	// send all log.* output to file
	log.SetOutput(logfile)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hostpath := r.Host
		ip := r.RemoteAddr
		log.Println("application host and ip details:", r.Method, ip)
		fmt.Fprintf(w, "Hello, World! You've reached %s\n", hostpath)
	})
	http.Handle("/metrics", http.HandlerFunc(metrics))
	log.Println("App started and listening on :8000")
	http.ListenAndServe(":8000", nil)
}

func metrics(w http.ResponseWriter, r *http.Request) {
	log.Println("metrics method received")
	fmt.Fprintf(w, "metrics endpoint reached\n")
}
