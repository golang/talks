// +build ignore,OMIT

// realmain runs the Subscribe example with a real RSS fetcher.
package main

import (
	"fmt"
	"math/rand"
	"time"

	rss "github.com/jteeuwen/go-pkg-rss"
)

// STARTITEM OMIT
// An Item is a stripped-down RSS item.
type Item struct{ Title, Channel, GUID string }

// STOPITEM OMIT

// STARTFETCHER OMIT
// A Fetcher fetches Items and returns the time when the next fetch should be
// attempted.  On failure, Fetch returns a non-nil error.
type Fetcher interface {
	Fetch() (items []Item, next time.Time, err error)
}

// STOPFETCHER OMIT

// STARTSUBSCRIPTION OMIT
// A Subscription delivers Items over a channel.  Close cancels the
// subscription, closes the Updates channel, and returns the last fetch error,
// if any.
type Subscription interface {
	Updates() <-chan Item
	Close() error
}

// STOPSUBSCRIPTION OMIT

// STARTSUBSCRIBE OMIT
// Subscribe returns a new Subscription that uses fetcher to fetch Items.
func Subscribe(fetcher Fetcher) Subscription {
	s := &sub{
		fetcher: fetcher,
		updates: make(chan Item),       // for Updates
		closing: make(chan chan error), // for Close
	}
	go s.loop()
	return s
}

// STOPSUBSCRIBE OMIT

// sub implements the Subscription interface.
type sub struct {
	fetcher Fetcher         // fetches items
	updates chan Item       // sends items to the user
	closing chan chan error // for Close
}

// STARTUPDATES OMIT
func (s *sub) Updates() <-chan Item {
	return s.updates
}

// STOPUPDATES OMIT

// STARTCLOSE OMIT
// STARTCLOSESIG OMIT
func (s *sub) Close() error {
	// STOPCLOSESIG OMIT
	errc := make(chan error)
	s.closing <- errc // HLchan
	return <-errc     // HLchan
}

// STOPCLOSE OMIT

// loopCloseOnly is a version of loop that includes only the logic
// that handles Close.
func (s *sub) loopCloseOnly() {
	// STARTCLOSEONLY OMIT
	var err error // set when Fetch fails
	for {
		select {
		case errc := <-s.closing: // HLchan
			errc <- err      // HLchan
			close(s.updates) // tells receiver we're done
			return
		}
	}
	// STOPCLOSEONLY OMIT
}

// loopFetchOnly is a version of loop that includes only the logic
// that calls Fetch.
func (s *sub) loopFetchOnly() {
	// STARTFETCHONLY OMIT
	var pending []Item // appended by fetch; consumed by send
	var next time.Time // initially January 1, year 0
	var err error
	for {
		var fetchDelay time.Duration // initally 0 (no delay)
		if now := time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
		}
		startFetch := time.After(fetchDelay)

		select {
		case <-startFetch:
			var fetched []Item
			fetched, next, err = s.fetcher.Fetch()
			if err != nil {
				next = time.Now().Add(10 * time.Second)
				break
			}
			pending = append(pending, fetched...)
		}
	}
	// STOPFETCHONLY OMIT
}

// loopSendOnly is a version of loop that includes only the logic for
// sending items to s.updates.
func (s *sub) loopSendOnly() {
	// STARTSENDONLY OMIT
	var pending []Item // appended by fetch; consumed by send
	for {
		var first Item
		var updates chan Item // HLupdates
		if len(pending) > 0 {
			first = pending[0]
			updates = s.updates // enable send case // HLupdates
		}

		select {
		case updates <- first:
			pending = pending[1:]
		}
	}
	// STOPSENDONLY OMIT
}

// mergedLoop is a version of loop that combines loopCloseOnly,
// loopFetchOnly, and loopSendOnly.
func (s *sub) mergedLoop() {
	// STARTFETCHVARS OMIT
	var pending []Item
	var next time.Time
	var err error
	// STOPFETCHVARS OMIT
	for {
		// STARTNOCAP OMIT
		var fetchDelay time.Duration
		if now := time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
		}
		startFetch := time.After(fetchDelay)
		// STOPNOCAP OMIT
		var first Item
		var updates chan Item
		if len(pending) > 0 {
			first = pending[0]
			updates = s.updates // enable send case
		}

		// STARTSELECT OMIT
		select {
		case errc := <-s.closing: // HLcases
			errc <- err
			close(s.updates)
			return
			// STARTFETCHCASE OMIT
		case <-startFetch: // HLcases
			var fetched []Item
			fetched, next, err = s.fetcher.Fetch() // HLfetch
			if err != nil {
				next = time.Now().Add(10 * time.Second)
				break
			}
			pending = append(pending, fetched...) // HLfetch
			// STOPFETCHCASE OMIT
		case updates <- first: // HLcases
			pending = pending[1:]
		}
		// STOPSELECT OMIT
	}
}

// dedupeLoop extends mergedLoop with deduping of fetched items.
func (s *sub) dedupeLoop() {
	const maxPending = 10
	// STARTSEEN OMIT
	var pending []Item
	var next time.Time
	var err error
	var seen = make(map[string]bool) // set of item.GUIDs // HLseen
	// STOPSEEN OMIT
	for {
		// STARTCAP OMIT
		var fetchDelay time.Duration
		if now := time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
		}
		var startFetch <-chan time.Time // HLcap
		if len(pending) < maxPending {  // HLcap
			startFetch = time.After(fetchDelay) // enable fetch case  // HLcap
		} // HLcap
		// STOPCAP OMIT
		var first Item
		var updates chan Item
		if len(pending) > 0 {
			first = pending[0]
			updates = s.updates // enable send case
		}
		select {
		case errc := <-s.closing:
			errc <- err
			close(s.updates)
			return
		// STARTDEDUPE OMIT
		case <-startFetch:
			var fetched []Item
			fetched, next, err = s.fetcher.Fetch() // HLfetch
			if err != nil {
				next = time.Now().Add(10 * time.Second)
				break
			}
			for _, item := range fetched {
				if !seen[item.GUID] { // HLdupe
					pending = append(pending, item) // HLdupe
					seen[item.GUID] = true          // HLdupe
				} // HLdupe
			}
			// STOPDEDUPE OMIT
		case updates <- first:
			pending = pending[1:]
		}
	}
}

// loop periodically fecthes Items, sends them on s.updates, and exits
// when Close is called.  It extends dedupeLoop with logic to run
// Fetch asynchronously.
func (s *sub) loop() {
	const maxPending = 10
	type fetchResult struct {
		fetched []Item
		next    time.Time
		err     error
	}
	// STARTFETCHDONE OMIT
	var fetchDone chan fetchResult // if non-nil, Fetch is running // HL
	// STOPFETCHDONE OMIT
	var pending []Item
	var next time.Time
	var err error
	var seen = make(map[string]bool)
	for {
		var fetchDelay time.Duration
		if now := time.Now(); next.After(now) {
			fetchDelay = next.Sub(now)
		}
		// STARTFETCHIF OMIT
		var startFetch <-chan time.Time
		if fetchDone == nil && len(pending) < maxPending { // HLfetch
			startFetch = time.After(fetchDelay) // enable fetch case
		}
		// STOPFETCHIF OMIT
		var first Item
		var updates chan Item
		if len(pending) > 0 {
			first = pending[0]
			updates = s.updates // enable send case
		}
		// STARTFETCHASYNC OMIT
		select {
		case <-startFetch: // HLfetch
			fetchDone = make(chan fetchResult, 1) // HLfetch
			go func() {
				fetched, next, err := s.fetcher.Fetch()
				fetchDone <- fetchResult{fetched, next, err}
			}()
		case result := <-fetchDone: // HLfetch
			fetchDone = nil // HLfetch
			// Use result.fetched, result.next, result.err
			// STOPFETCHASYNC OMIT
			fetched := result.fetched
			next, err = result.next, result.err
			if err != nil {
				next = time.Now().Add(10 * time.Second)
				break
			}
			for _, item := range fetched {
				if id := item.GUID; !seen[id] { // HLdupe
					pending = append(pending, item)
					seen[id] = true // HLdupe
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

// naiveMerge is a version of Merge that doesn't quite work right.  In
// particular, the goroutines it starts may block forever on m.updates
// if the receiver stops receiving.
type naiveMerge struct {
	subs    []Subscription
	updates chan Item
}

// STARTNAIVEMERGE OMIT
func NaiveMerge(subs ...Subscription) Subscription {
	m := &naiveMerge{
		subs:    subs,
		updates: make(chan Item),
	}
	// STARTNAIVEMERGELOOP OMIT
	for _, sub := range subs {
		go func(s Subscription) {
			for it := range s.Updates() {
				m.updates <- it // HL
			}
		}(sub)
	}
	// STOPNAIVEMERGELOOP OMIT
	return m
}

// STOPNAIVEMERGE OMIT

// STARTNAIVEMERGECLOSE OMIT
func (m *naiveMerge) Close() (err error) {
	for _, sub := range m.subs {
		if e := sub.Close(); err == nil && e != nil {
			err = e
		}
	}
	close(m.updates) // HL
	return
}

// STOPNAIVEMERGECLOSE OMIT

func (m *naiveMerge) Updates() <-chan Item {
	return m.updates
}

type merge struct {
	subs    []Subscription
	updates chan Item
	quit    chan struct{}
	errs    chan error
}

// STARTMERGESIG OMIT
// Merge returns a Subscription that merges the item streams from subs.
// Closing the merged subscription closes subs.
func Merge(subs ...Subscription) Subscription {
	// STOPMERGESIG OMIT
	m := &merge{
		subs:    subs,
		updates: make(chan Item),
		quit:    make(chan struct{}),
		errs:    make(chan error),
	}
	// STARTMERGE OMIT
	for _, sub := range subs {
		go func(s Subscription) {
			for {
				var it Item
				select {
				case it = <-s.Updates():
				case <-m.quit: // HL
					m.errs <- s.Close() // HL
					return              // HL
				}
				select {
				case m.updates <- it:
				case <-m.quit: // HL
					m.errs <- s.Close() // HL
					return              // HL
				}
			}
		}(sub)
	}
	// STOPMERGE OMIT
	return m
}

func (m *merge) Updates() <-chan Item {
	return m.updates
}

// STARTMERGECLOSE OMIT
func (m *merge) Close() (err error) {
	close(m.quit) // HL
	for _ = range m.subs {
		if e := <-m.errs; e != nil { // HL
			err = e
		}
	}
	close(m.updates) // HL
	return
}

// STOPMERGECLOSE OMIT

// NaiveDedupe converts a stream of Items that may contain duplicates
// into one that doesn't.
func NaiveDedupe(in <-chan Item) <-chan Item {
	out := make(chan Item)
	go func() {
		seen := make(map[string]bool)
		for it := range in {
			if !seen[it.GUID] {
				// BUG: this send blocks if the
				// receiver closes the Subscription
				// and stops receiving.
				out <- it // HL
				seen[it.GUID] = true
			}
		}
		close(out)
	}()
	return out
}

type deduper struct {
	s       Subscription
	updates chan Item
	closing chan chan error
}

// Dedupe converts a Subscription that may send duplicate Items into
// one that doesn't.
func Dedupe(s Subscription) Subscription {
	d := &deduper{
		s:       s,
		updates: make(chan Item),
		closing: make(chan chan error),
	}
	go d.loop()
	return d
}

func (d *deduper) loop() {
	in := d.s.Updates() // enable receive
	var pending Item
	var out chan Item // disable send
	seen := make(map[string]bool)
	for {
		select {
		case it := <-in:
			if !seen[it.GUID] {
				pending = it
				in = nil        // disable receive
				out = d.updates // enable send
				seen[it.GUID] = true
			}
		case out <- pending:
			in = d.s.Updates() // enable receive
			out = nil          // disable send
		case errc := <-d.closing:
			err := d.s.Close()
			errc <- err
			close(d.updates)
			return
		}
	}
}

func (d *deduper) Close() error {
	errc := make(chan error)
	d.closing <- errc
	return <-errc
}

func (d *deduper) Updates() <-chan Item {
	return d.updates
}

// FakeFetch causes Fetch to use a fake fetcher instead of the real
// one.
var FakeFetch bool

// Fetch returns a Fetcher for Items from domain.
func Fetch(domain string) Fetcher {
	if FakeFetch {
		return fakeFetch(domain)
	}
	return realFetch(domain)
}

func fakeFetch(domain string) Fetcher {
	return &fakeFetcher{channel: domain}
}

type fakeFetcher struct {
	channel string
	items   []Item
}

// FakeDuplicates causes the fake fetcher to return duplicate items.
var FakeDuplicates bool

func (f *fakeFetcher) Fetch() (items []Item, next time.Time, err error) {
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

// realFetch returns a fetcher for the specified blogger domain.
func realFetch(domain string) Fetcher {
	return NewFetcher(fmt.Sprintf("http://%s/feeds/posts/default?alt=rss", domain))
}

type fetcher struct {
	uri   string
	feed  *rss.Feed
	items []Item
}

// NewFetcher returns a Fetcher for uri.
func NewFetcher(uri string) Fetcher {
	f := &fetcher{
		uri: uri,
	}
	newChans := func(feed *rss.Feed, chans []*rss.Channel) {}
	newItems := func(feed *rss.Feed, ch *rss.Channel, items []*rss.Item) {
		for _, it := range items {
			f.items = append(f.items, Item{
				Channel: ch.Title,
				GUID:    it.Guid,
				Title:   it.Title,
			})
		}
	}
	f.feed = rss.New(1 /*minimum interval in minutes*/, true /*respect limit*/, newChans, newItems)
	return f
}

func (f *fetcher) Fetch() (items []Item, next time.Time, err error) {
	fmt.Println("fetching", f.uri)
	if err = f.feed.Fetch(f.uri, nil); err != nil {
		return
	}
	items = f.items
	f.items = nil
	next = time.Now().Add(time.Duration(f.feed.SecondsTillUpdate()) * time.Second)
	return
}

// TODO: in a longer talk: move the Subscribe function onto a Reader type, to
// support dynamically adding and removing Subscriptions.  Reader should dedupe.

// TODO: in a longer talk: make successive Subscribe calls for the same uri
// share the same underlying Subscription, but provide duplicate streams.

func init() {
	rand.Seed(time.Now().UnixNano())
}

// STARTMAIN OMIT
func main() {
	// STARTMERGECALL OMIT
	// Subscribe to some feeds, and create a merged update stream.
	merged := Merge(
		Subscribe(Fetch("blog.golang.org")),
		Subscribe(Fetch("googleblog.blogspot.com")),
		Subscribe(Fetch("googledevelopers.blogspot.com")))
	// STOPMERGECALL OMIT

	// Close the subscriptions after some time.
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("closed:", merged.Close())
	})

	// Print the stream.
	for it := range merged.Updates() {
		fmt.Println(it.Channel, it.Title)
	}

	// Uncomment the panic below to dump the stack traces.  This
	// will show several stacks for persistent HTTP connections
	// created by the real RSS client.  To clean these up, we'll
	// need to extend Fetcher with a Close method and plumb this
	// through the RSS client implementation.
	//
	// panic("show me the stacks")
}

// STOPMAIN OMIT
