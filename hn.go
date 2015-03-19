package main

import (
	"fmt"
	hn "github.com/munrocape/hn/client"
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
	if timeSinceLast > timeToExpire {
		return true
	}
	return false
}

func generateNewResponse() (string, error) {
	c := getClient()
	var stories []int
	stories, err := c.GetTopStories(10)
	if err != nil {
		return "", err
	}

	var urls [11]string
	urls[0] = "Top Stories from Hacker News"
	for index, element := range stories[:10] {
		index = index + 1
		story, err := c.GetStory(element)
		if err == nil {
			urls[index] = fmt.Sprintf("<%s|%d. %s> - [<https://news.ycombinator.com/item?id=%d|%d comments>]", story.Url, index, story.Title, element, story.Descendants)
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
