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
	getHnTop10()
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
	switch {
	case news_source == "hn":
		stories, err := getHnTop10()
		if err == nil {
			w.Write([]byte(stories))
		} else {
			w.Write([]byte("Server Error - Firebase could not be reached"))
		}
		return
	}
	w.Write([]byte("can't find that one! " + news_source))
}