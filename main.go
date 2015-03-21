package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

var (
	timeToExpire time.Duration
)

func main() {
	// Parse port from command-line parameters
	port := flag.String("port", "8080", "HTTP Port to listen on")
	flag.Parse()

	// declare variables
	timeToExpire = 10 * time.Minute

	// Start our Server
	log.Println("Starting Server on", *port)
	http.HandleFunc("/", index)
	http.HandleFunc("/news", news)
	log.Fatal(http.ListenAndServe(":"+*port, nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Slacker News, a slack integration to provide current news. Check it out at www.github.com/munrocape/slacker-news"))
}

func news(w http.ResponseWriter, r *http.Request) {
	news_source := r.URL.Query().Get("text")
	log.Println(news_source)
	tokens := strings.Split(news_source, " ")
	var source, argument string
	if len(tokens) == 2 {
		source, argument = tokens[0], tokens[1]
	} else {
		source, argument = tokens[0], ""
	}
	switch {
	case source == "hn":
		stories, err := getHnTop10()
		if err == nil {
			w.Write([]byte(stories))
		} else {
			w.Write([]byte("Server Error - Firebase could not be reached"))
		}
		return
	case source == "ph":
		posts, err := GetPhTop10()
		if err == nil {
			w.Write([]byte(posts))
		} else {
			w.Write([]byte("Server Error - Product Hunt could not be reached"))
		}
		return
	case source == "vice":
		articles, err := GetViceTop10()
		if err == nil {
			w.Write([]byte(articles))
		} else {
			w.Write([]byte("Server Error - Vice News could not be reached"))
		}
		return
	case source == "bbc":
		articles, err := GetBbcTop10(argument)
		if err == nil {
			w.Write([]byte(articles))
		} else {
			if strings.Contains(err.Error(), "Invalid feed selection") {
				response := fmt.Sprintf("That is an invalid BBC category: %s\nTry `/news list_sources` to view all sources.", argument)
				w.Write([]byte(response))
			} else {
				w.Write([]byte("Server Error - the BBC could not be reached"))
			}
		}
		return
	case source == "538":
		articles, err := GetFteTop10(argument)
		if err == nil {
			w.Write([]byte(articles))
		} else {
			if strings.Contains(err.Error(), "Invalid feed selection") {
				response := fmt.Sprintf("That is an invalid FiveThirtyEight category: %s\nTry `/news list_sources` to view all sources.", argument)
				w.Write([]byte(response))
			} else {
				w.Write([]byte("Server Error - FiveThirtyEight could not be reached"))
			}
		}
		return
	case source == "list_sources":
		w.Write([]byte(GetSources()))
		return
	}
	user_argument := fmt.Sprintf("%s %s", source, argument)
	w.Write([]byte("Hmm.. I can't figure out what news you are looking for :( I received \"" + strings.TrimSpace(user_argument) + "\"\nTry `/news list_sources` to view all sources."))
}

func GetSources() string {
	hn := "Hacker News: hn\n"
	ph := "Product Hunt: ph\n"
	vice := "Vice News: vice\n"
	fte := "FiveThirtyEight: 538 [" + GetFteSources() + "]\n"
	bbc := "BBC: bbc [" + GetBbcSources() + "]\n"
	return fmt.Sprintf("%s%s%s%s%s", hn, ph, vice, fte, bbc)
}
