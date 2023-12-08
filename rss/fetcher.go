package rss

import (
	"fmt"
	"math/rand"
	"time"
)

// A Fetcher fetches Items and returns the time when the next fetch should be
// attempted.  On failure, Fetch returns a non-nil error.
type Fetcher interface {
	Fetch() (items []Item, next time.Time, err error)
}

type fetcher struct {
	channel string
	items   []Item
}

// Fetch returns a Fetcher for uri.
func Fetch(domain string) Fetcher {
	return &fetcher{channel: domain}
}

// FakeDuplicates causes the fake fetcher to return duplicate items.
var FakeDuplicates bool

func (f *fetcher) Fetch() (items []Item, next time.Time, err error) {
	now := time.Now()
	next = now.Add(time.Duration(rand.Intn(5)) * 500 * time.Millisecond)
	item := Item{
		Channel: f.channel,
		Title:   fmt.Sprintf("Item %d", len(f.items)),
	}
	item.GUID = item.Channel + "/" + item.Title
	f.items = append(f.items, item)
	if FakeDuplicates {
		items = f.items
	} else {
		items = []Item{item}
	}
	return
}
