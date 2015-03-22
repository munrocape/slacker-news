package main

import (
	"fmt"
	bbc "github.com/munrocape/bbc/bbcclient"
	"strings"
	"time"
)

var (
	BbcClient           *bbc.Client
	currentBbcResponses map[string]string
	bbcTimestamps       map[string]time.Time
)

func GetBbcTop10(category string) (string, error) {
	var err error
	var rep string
	if ExpiredResponse(bbcTimestamps[category]) {
		rep, err = generateNewBbcResponse(category)
		currentBbcResponses[category] = rep
	}
	return currentBbcResponses[category], err
}

func generateNewBbcResponse(category string) (string, error) {
	c := getBbcClient()
	rss, err := c.GetFeed(category)
	if err != nil {
		return "", err
	}

	var urls [11]string
	var url = c.GetUrl(category)
	urls[0] = fmt.Sprintf("Top Stories from <%s|BBC %s>", url, c.GetPretty(category))
	items := rss.Channel.Items
	for index, element := range items {
		index = index + 1
		if index < 11 {
			urls[index] = fmt.Sprintf("%d. <%s|%s>", index, element.Link, element.Title)
		}
	}

	response := strings.Join(urls[:], "\n")
	bbcTimestamps[category] = time.Now().Local()
	return response, nil
}

func GetBbcSources() string {
	c := getBbcClient()
	res := ""
	first := true
	for k, _ := range c.NewsCategories {
		if first {
			res = res + k
			first = false
		} else {
			res = res + ", " + k
		}

	}
	for k, _ := range c.SportsCategories {
		res = res + ", " + k
	}
	return res
}

func getBbcClient() *bbc.Client {
	if BbcClient == nil {
		BbcClient = bbc.NewClient()
		currentBbcResponses = make(map[string]string)
		bbcTimestamps = make(map[string]time.Time)
	}
	return BbcClient
}
