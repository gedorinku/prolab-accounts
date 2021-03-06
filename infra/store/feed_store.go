package store

import (
	"github.com/mmcdole/gofeed"
)

// FeedStore provides feed
type FeedStore interface {
	GetFeedURL(url string) (string, error)
	IsValidFeedURL(url string) error
	GetFeed(feedURL string) (*gofeed.Feed, error)
}
