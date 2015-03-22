package main

import (
	"fmt"
	ph "github.com/munrocape/ph/phclient"
	"strings"
	"time"
)

var (
	PhClient           *ph.Client
	currentPhResponse  string
	currentPhTimestamp time.Time
)

func GetPhTop10() (string, error) {
	var err error
	if ExpiredResponse(currentPhTimestamp) {
		currentPhResponse, err = generateNewPhResponse()
	}
	return currentPhResponse, err
}

func generateNewPhResponse() (string, error) {
	c := getPhClient()
	var posts ph.PostsResponse
	posts, err := c.GetPostsToday()
	if err != nil {
		return "", err
	}

	var urls [11]string
	urls[0] = "Top Stories from <www.producthunt.com|Product Hunt>"
	for index, element := range posts.Posts {
		index = index + 1
		if index < 11 {
			if err == nil {
				urls[index] = fmt.Sprintf("%d. <%s|%s>: %s - [<%s|%d comments>]", index, element.RedirectUrl, element.Name, element.Tagline, element.DiscussionUrl, element.CommentsCount)
			} else {
				urls[index] = "Server Error - Firebase did not return the submission information."
			}
		}
	}

	response := strings.Join(urls[:], "\n")
	currentPhTimestamp = time.Now().Local()
	return response, nil
}

func getPhClient() *ph.Client {
	if PhClient == nil {
		PhClient = ph.NewClient()
	}
	return PhClient
}
