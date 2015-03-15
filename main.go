package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Parse port from command-line parameters
	port := flag.String("port", "8080", "HTTP Port to listen on")
	flag.Parse()

	// Start our Server
	log.Println("Starting Server on", *port)
	http.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Slacker News, a slack integration. Check it out at www.github.com/munrocape/slacker-news"))
}