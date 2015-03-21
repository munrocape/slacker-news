package main

import (
	"fmt"
	vice "github.com/munrocape/vice/viceclient"
	"strings"
	"time"
)

var (
	ViceClient           *vice.Client
	currentViceResponse  string
	currentViceTimestamp time.Time
)

func GetViceTop10() (string, error) {
	var err error
	if expiredViceResponse() {
		currentViceResponse, err = generateNewViceResponse()
	}
	return currentViceResponse, err
}

func expiredViceResponse() bool {
	timeSinceLast := time.Since(currentViceTimestamp)
	if timeSinceLast > timeToExpire {
		return true
	}
	return false
}

func generateNewViceResponse() (string, error) {
	c := getViceClient()
	var rss vice.Rss
	rss, err := c.GetFeed()
	if err != nil {
		return "", err
	}

	var urls [11]string
	urls[0] = "Top Stories from <https://www.vice.com|Vice News>"
	items := rss.Channel.Items
	for index, element := range items {
		index = index + 1
		if index < 11 {
			urls[index] = fmt.Sprintf("<%s|%d. %s>", element.Link, index, element.Title)

		}
	}

	response := strings.Join(urls[:], "\n")
	currentViceTimestamp = time.Now().Local()
	return response, nil
}

func getViceClient() *vice.Client {
	if ViceClient == nil {
		ViceClient = vice.NewClient()
	}
	return ViceClient
}
