package rss

import "time"

// StartPolling periodically fetches Items, sends them on s.updates, and exits
// when Close is called.  It extends dedupeLoop with logic to run
// Fetch asynchronously.
func (s *sub) StartPolling() {
	const maxPending = 10
	type fetchResult struct {
		fetched []Item
		next    time.Time
		err     error
	}
	var fetchDone chan fetchResult // if non-nil, Fetch is running // HL
	var pending []Item
	var next time.Time
	var err error
	var seen = make(map[string]bool)
	for {
		var fetchDelay time.Duration
		if now := time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
		}
		var startFetch <-chan time.Time
		if fetchDone == nil && len(pending) < maxPending {
			startFetch = time.After(fetchDelay) // enable fetch case
		}
		var first Item
		var updates chan Item
		if len(pending) > 0 {
			first = pending[0]
			updates = s.updates // enable send case
		}
		select {
		case <-startFetch:
			fetchDone = make(chan fetchResult, 1)
			go func() {
				fetched, next, err := s.fetcher.Fetch()
				fetchDone <- fetchResult{fetched, next, err}
			}()
		case result := <-fetchDone:
			fetchDone = nil
			// Use result.fetched, result.next, result.err
			fetched := result.fetched
			next, err = result.next, result.err
			if err != nil {
				next = time.Now().Add(10 * time.Second)
				break
			}
			for _, item := range fetched {
				if id := item.GUID; !seen[id] {
					pending = append(pending, item)
					seen[id] = true
				}
			}
		case errc := <-s.closing:
			errc <- err
			close(s.updates)
			return
		case updates <- first:
			pending = pending[1:]
		}
	}
}
