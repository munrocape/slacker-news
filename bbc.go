package main

import (
	"fmt"
	bbc  "github.com/munrocape/bbc/bbcclient"
	"strings"
	"time"
)

var (
	BbcClient           *bbc.Client
	currentBbcResponse  string
	currentBbcTimestamp time.Time
)

func getBbcTop10(category string) (string, error) {
	var err error
	if expiredBbcResponse() {
		currentBbcResponse, err = generateNewBbcResponse(category)
	}
	return currentBbcResponse, err
}

func expiredBbcResponse() bool {
	timeSinceLast := time.Since(currentBbcTimestamp)
	if timeSinceLast > timeToExpire {
		return true
	}
	return false
}

func generateNewBbcResponse(category string) (string, error) {
	c := getBbcClient()
	rss, err := c.GetFeed(category)
	if err != nil {
		return "", err
	}

	var urls [11]string
	urls[0] = "Top Stories from BBC " + category
	items := rss.Channel.Items
	for index, element := range items {
		index = index + 1
		if (index < 11){
			urls[index] = fmt.Sprintf("<%s|%d. %s>\n\t%s", element.Link, index, element.Title, element.Description)
		}
	}

	response := strings.Join(urls[:], "\n")
	currentBbcTimestamp = time.Now().Local()
	return response, nil
}

func GetBbcSources() string {
	c := getBbcClient()
	res := ""
	for k, _ := range c.NewsCategories {
		res = res + " " + k
	}
	for k, _ := range c.SportsCategories {
		res = res + " " + k
	}
	return res
}

func getBbcClient() *bbc.Client {
	if BbcClient == nil {
		BbcClient = bbc.NewClient()
	}
	return BbcClient
}
