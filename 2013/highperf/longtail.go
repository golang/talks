// +build OMIT

package pkg

// long_tail_memcache_bad
func myHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	// ...
	// regular request handling
	// ...

	go memcache.Set(c, &memcache.Item{
		Key:   key,
		Value: data,
	})
}

// long_tail_memcache_good
func myHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	// ...
	// regular request handling
	// ...

	// Save to memcache, but only wait up to 3ms.
	done := make(chan bool, 1) // NB: buffered
	go func() {
		memcache.Set(c, &memcache.Item{
			Key:   key,
			Value: data,
		})
		done <- true
	}()
	select { // HL
	case <-done: // HL
	case <-time.After(3 * time.Millisecond): // HL
	} // HL
}
