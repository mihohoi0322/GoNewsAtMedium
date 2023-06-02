package main

import (
	"github.com/mmcdole/gofeed"
	"testing"
)

func GetFeedItems(url string) ([]*gofeed.Item, error) {
	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(url)
	if err != nil {
		return nil, err
	}
	return feed.Items, nil
}

func TestGetFeedItems(t *testing.T) {
	items, err := GetFeedItems("https://medium.com/feed/tag/go")
	if err != nil {
		t.Errorf("failed to get RSS feed: %v", err)
	}
	if len(items) == 0 {
		t.Errorf("no items found in the feed")
	}
	for _, item := range items {
		if item.Title == "" {
			t.Errorf("found item with empty title")
		}
	}
}
