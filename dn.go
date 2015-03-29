package main

import (
	"fmt"
	dn "github.com/munrocape/dn/dnclient"
	"strings"
	"time"
)

var (
	DnClient           *dn.Client
	currentDnStories  string
	currentDnMotd string
	currentDnStoriesTimestamp time.Time
	currentDnMotdTimestamp time.Time
)

func GetDnArgument(argument string) (string, error){
	if (argument == "motd"){
		return GetDnMotd()
	} else if (argument == "news"){
		return GetDnTop10()
	} else {
		return "", fmt.Errorf("Invalid argument")
	}
}

func GetDnTop10() (string, error) {
	var err error
	if ExpiredResponse(currentDnStoriesTimestamp) {
		currentDnStories, err = generateNewDnResponse()
	}
	return currentDnStories, err
}

func GetDnMotd() (string, error){
	var err error
	if ExpiredResponse(currentDnMotdTimestamp) {
		currentDnMotd, err = generateNewDnMotd()
	}
	return currentDnMotd, err
}

func generateNewDnMotd() (string, error) {
	c := getDnClient()
	var motd dn.Motd
	motd, err := c.GetMotd()
	if err != nil {
		return "", err
	}

	var response [2]string
	response[0] = "The Message of the Day from <www.news.layervault.com|Designer News>"
	response[1] = fmt.Sprintf("%s (%d votes) - submitted by %s", motd.Message, motd.UpvoteCount, motd.UserDisplayName)
	final := strings.Join(response[:], "\n")
	currentDnMotdTimestamp = time.Now().Local()
	return final, nil
}

func generateNewDnResponse() (string, error) {
	c := getDnClient()
	var storyWrapper dn.Stories
	storyWrapper, err := c.GetStories()
	stories := storyWrapper.Stories

	if err != nil {
		return "", err
	}

	count := len(stories)
	var urls [11]string
	urls[0] = "Top Stories from <www.news.layervault.com|Designer News>"
	for index, element := range stories[:count] {
		if (index < 10){
			index = index + 1
			if err == nil {
				urls[index] = fmt.Sprintf("%d. <%s|%s> - [<%s|%d comments>]", index, element.Url, element.Title, element.SiteUrl, element.CommentCount)
			} else {
				urls[index] = "Server Error - Designer News did not return the story information."
			}
		}
	}

	response := strings.Join(urls[:], "\n")
	currentDnStoriesTimestamp = time.Now().Local()
	return response, nil
}

func getDnClient() *dn.Client {
	if DnClient == nil {
		DnClient = dn.NewClient()
	}
	return DnClient
}
