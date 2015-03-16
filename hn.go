package main

import (
	"fmt"
	hn "github.com/caser/gophernews"
	"strings"
	"time"
)

var (
	client           *hn.Client
	currentResponse  string
	currentTimestamp time.Time
)

func getHnTop10() (string, error) {
	var err error
	if expiredResponse() {
		currentResponse, err = generateNewResponse()
	}
	return currentResponse, err
}

func expiredResponse() bool {
	timeSinceLast := time.Since(currentTimestamp)
	if timeSinceLast > (10 * time.Minute) {
		return true
	}
	return false
}

func generateNewResponse() (string, error) {
	c := getClient()
	var stories []int
	stories, err := c.GetTop100()
	if err != nil {
		return "", err
	}

	var urls [11]string
	urls[0] = "Top Stories from Hacker News"
	for index, element := range stories[:10] {
		index = index + 1
		story, err := fetchStory(element)
		if err == nil {
			urls[index] = fmt.Sprintf("<%s|%d. %s>", story.URL, index+1, story.Title)
		} else {
			urls[index] = "Server Error - Firebase did not return the story information."
		}
	}
	
	response := strings.Join(urls[:], "\n")
	currentTimestamp = time.Now().Local()
	return response, nil
}

func getClient() *hn.Client {
	if client == nil {
		client = hn.NewClient()
	}
	return client
}

func fetchStory(element int) (hn.Story, error) {
	c := getClient()
	// for some reason, Firebase returns EOF on occasion
	// retry a few times in case this happens
	var err error
	var story hn.Story
	for i := 0; i < 5; i++ {
		story, err = c.GetStory(element)
		if err == nil {
			return story, err
		}
	}
	return *new(hn.Story), err
}
