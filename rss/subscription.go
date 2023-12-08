package rss

// An Item is a stripped-down RSS item.
type Item struct{ Title, Channel, GUID string }

// A Subscription delivers Items over a channel.  Close cancels the
// subscription, closes the Updates channel, and returns the last fetch error,
// if any.
type Subscription interface {
	Updates() <-chan Item
	Close() error
}

// Subscribe returns a new Subscription that uses fetcher to fetch Items.
func Subscribe(fetcher Fetcher) Subscription {
	s := &sub{
		fetcher: fetcher,
		updates: make(chan Item),       // for Updates
		closing: make(chan chan error), // for Closenext
	}
	go s.StartPolling()
	return s
}

// sub implements the Subscription interface.
type sub struct {
	fetcher Fetcher         // fetches items
	updates chan Item       // sends items to the user
	closing chan chan error // for Close
}

func (s *sub) Updates() <-chan Item {
	return s.updates
}

func (s *sub) Close() error {
	errc := make(chan error)
	s.closing <- errc
	return <-errc
}
