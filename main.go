package main

import (
	"fmt"
	"log"

	"github.com/mmcdole/gofeed"
)

func main() {
	fp := gofeed.NewParser()

	// Get the RSS feed of Go-related articles on Medium
	feed, err := fp.ParseURL("https://medium.com/feed/tag/go")
	if err != nil {
		log.Printf("Error while fetching the RSS feed: %v", err)
		return
	}

	// Print the titles, links, and publication time of the first 10 items
	for i, item := range feed.Items {
		if i >= 10 {
			break
		}
		fmt.Printf("Article %d:\n", i+1)
		fmt.Printf("  Title: %s\n", item.Title)
		fmt.Printf("  Link: %s\n", item.Link)
		if item.PublishedParsed != nil {
			fmt.Printf("  Published at: %s\n", item.PublishedParsed.Format("2006/01/02"))
		} else {
			fmt.Println("  Published at: unknown")
		}
		fmt.Println()
	}
}
