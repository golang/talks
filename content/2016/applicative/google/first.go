package google

import (
	"errors"
	"time"
)

func First(replicas ...SearchFunc) SearchFunc { // HL
	return func(query string) Result {
		c := make(chan Result, len(replicas))
		searchReplica := func(i int) {
			c <- replicas[i](query)
		}
		for i := range replicas {
			go searchReplica(i) // HL
		}
		return <-c
	}
}

// START OMIT
var (
	replicatedWeb   = First(Web1, Web2)     // HL
	replicatedImage = First(Image1, Image2) // HL
	replicatedVideo = First(Video1, Video2) // HL
)

func SearchReplicated(query string, timeout time.Duration) ([]Result, error) {
	timer := time.After(timeout)
	c := make(chan Result, 3)
	go func() { c <- replicatedWeb(query) }()   // HL
	go func() { c <- replicatedImage(query) }() // HL
	go func() { c <- replicatedVideo(query) }() // HL
	// STOP OMIT

	var results []Result
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timer:
			return results, errors.New("timed out")
		}
	}
	return results, nil
}
