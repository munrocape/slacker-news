package main

import(
	hn "github.com/caser/gophernews"
	"fmt"
	"strings"
)

var (
	client *hn.Client
)

func getClient() *hn.Client {
	if client == nil {
		client = hn.NewClient()
	}
	return client
}

func getHnTop10() (string, error) { 
	c := getClient()
	var stories []int
	stories, err := c.GetTop100()
	if err != nil {
		return "", err
	}
	
	var urls [10]string
	for index, element := range stories[:10] {
		story, err := fetchStory(element)
		if err == nil {
			urls[index] = fmt.Sprintf("<%s|%s>", story.URL, story.Title)
		} else {
			urls[index] = "Server Error - Firebase did not return the story information."
		}
	}
	response := strings.Join(urls[:], "\n")
	return response, nil
}

func fetchStory(element int) (hn.Story, error) {
	c := getClient()
	// for some reason, Firebase returns EOF on occasion
	// retry a few times in case this happens
	var err error
	for i := 0; i < 5; i++ {
	    story, err := c.GetStory(element)
	    if err == nil {
	    	return story, err
	    }
	}
	return *new(hn.Story), err
}