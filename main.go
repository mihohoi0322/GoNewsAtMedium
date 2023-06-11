package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"os"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()
	TOKEN := os.Getenv("SLACK_TOKEN")
	CHANNEL_ID := os.Getenv("SLACK_CHANNEL_ID")

	// Get the RSS feed of Go-related articles on Medium
	feed, err := fp.ParseURL("https://medium.com/feed/tag/go")
	if err != nil {
		log.Printf("Error while fetching the RSS feed: %v", err)
		return
	}

	// Create a new Slack client with your token
	slackClient := slack.New(TOKEN)
	if slackClient == nil {
		log.Printf("Error while creating a Slack client")
		return
	}

	// Print the titles, links, and publication time of the first 10 items
	for i, item := range feed.Items {
		if i >= 5 {
			break
		}
		msg := fmt.Sprintf("Article %d:\nTitle: %s\nLink: %s\n", i+1, item.Title, item.Link)
		if item.PublishedParsed != nil {
			msg += fmt.Sprintf("Published at: %s\n", item.PublishedParsed.Format("2006/01/02"))
		} else {
			msg += "Published at: unknown\n"
		}

		// Send each article to a specific channel
		_, _, err := slackClient.PostMessage(
			CHANNEL_ID,
			slack.MsgOptionText(msg, false),
		)
		if err != nil {
			log.Printf("Error while sending a message to Slack: %v", err)
			return
		}
	}

	fmt.Println("RSS feed has been fetched and sent to Slack successfully!")
}
