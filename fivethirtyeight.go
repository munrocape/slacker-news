package main

import (
	"fmt"
	fte "github.com/munrocape/fivethirtyeight/fivethirtyeightclient"
	"strings"
	"time"
)

var (
	FteClient           *fte.Client
	currentFteResponses map[string]string
	fteTimestamps       map[string]time.Time
)

func GetFteTop10(category string) (string, error) {
	var err error
	var rep string
	if expiredFteResponse(fteTimestamps[category]) {
		rep, err = generateNewFteResponse(category)
		currentFteResponses[category] = rep
	}
	return currentFteResponses[category], err
}

func expiredFteResponse(timestamp time.Time) bool {
	timeSinceLast := time.Since(timestamp)
	if timeSinceLast > timeToExpire {
		return true
	}
	return false
}

func generateNewFteResponse(category string) (string, error) {
	c := getFteClient()
	rss, err := c.GetFeed(category)
	if err != nil {
		return "", err
	}

	var urls [11]string
	pretty := c.GetPretty(category)
	uri := c.GetUri(category)
	homepage := fmt.Sprintf("http://www.fivethirtyeight.com/%s", uri)
	title := fmt.Sprintf("Top Stories from <%s|Five Thirty Eight %s>", homepage, pretty)
	urls[0] = title
	items := rss.Channel.Items
	for index, element := range items {
		index = index + 1
		if index < 11 {
			urls[index] = fmt.Sprintf("%d. <%s|%s>", index, element.Link, element.Title)
		}
	}

	response := strings.Join(urls[:], "\n")
	fteTimestamps[category] = time.Now().Local()
	return response, nil
}

func GetFteSources() string {
	c := getFteClient()
	res := ""
	first := true
	for k, _ := range c.Categories {
		if first {
			res = res + k
			first = false
		} else {
			res = res + ", " + k
		}

	}
	return res
}

func getFteClient() *fte.Client {
	if FteClient == nil {
		FteClient = fte.NewClient()
		currentFteResponses = make(map[string]string)
		fteTimestamps = make(map[string]time.Time)
		initializeFteTimestampMap(FteClient)
		initializeFteResponseMap(FteClient)
	}
	return FteClient
}

func initializeFteTimestampMap(c *fte.Client) {
	for k, _ := range c.Categories {
		fteTimestamps[k] = time.Now().Local().AddDate(0, 0, -11)
	}
}

func initializeFteResponseMap(c *fte.Client) {
	for k, _ := range c.Categories {
		currentFteResponses[k] = ""
	}
}
