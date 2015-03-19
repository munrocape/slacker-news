package main

import (
	"fmt"
	hn "github.com/munrocape/hn/hnclient"
	"strings"
	"time"
)

var (
	HnClient           *hn.Client
	currentHnResponse  string
	currentHnTimestamp time.Time
)

func getHnTop10() (string, error) {
	var err error
	if expiredHnResponse() {
		currentHnResponse, err = generateNewHnResponse()
	}
	return currentHnResponse, err
}

func expiredHnResponse() bool {
	timeSinceLast := time.Since(currentHnTimestamp)
	if timeSinceLast > timeToExpire {
		return true
	}
	return false
}

func generateNewHnResponse() (string, error) {
	c := getHnClient()
	var stories []int
	count := 10
	stories, err := c.GetTopStories(count)
	if err != nil {
		return "", err
	}

	var urls [11]string
	urls[0] = "Top Stories from Hacker News"
	for index, element := range stories[:count] {
		index = index + 1
		story, err := c.GetStory(element)
		if err == nil {
			urls[index] = fmt.Sprintf("<%s|%d. %s> - [<https://news.ycombinator.com/item?id=%d|%d comments>]", story.Url, index, story.Title, element, story.Descendants)
		} else {
			urls[index] = "Server Error - Firebase did not return the story information."
		}
	}

	response := strings.Join(urls[:], "\n")
	currentHnTimestamp = time.Now().Local()
	return response, nil
}

func getHnClient() *hn.Client {
	if HnClient == nil {
		HnClient = hn.NewClient()
	}
	return HnClient
}
