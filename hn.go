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
	if ExpiredResponse(currentHnTimestamp) {
		currentHnResponse, err = generateNewHnResponse()
	}
	return currentHnResponse, err
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
	urls[0] = "Top Stories from <www.news.ycombinator.com|Hacker News>"
	for index, element := range stories[:count] {
		index = index + 1
		item, err := c.GetItem(element)

		if err == nil {
			if item.Type == "story" {
				if item.Url == "" {
					// It is an AskHN post
					urls[index] = fmt.Sprintf("%d. <https://news.ycombinator.com/item?id=%d|%s> - [<https://news.ycombinator.com/item?id=%d|%d comments>]", index, element, item.Title, element, item.Descendants)
				} else {
					urls[index] = fmt.Sprintf("%d. <%s|%s> - [<https://news.ycombinator.com/item?id=%d|%d comments>]", index, item.Url, item.Title, element, item.Descendants)
				}

			} else {
				urls[index] = fmt.Sprintf("%d. <https://news.ycombinator.com/item?id=%d|%s>", index, element, item.Title)
			}
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
