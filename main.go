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
	http.HandleFunc("/news", news)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Slacker News, a slack integration. Check it out at www.github.com/munrocape/slacker-news"))
}

func news(w http.ResponseWriter, r *http.Request) {
	news_source := r.URL.Query().Get("text")
	log.Println(news_source)
	w.Write([]byte("Getting stories for " + news_source))
}