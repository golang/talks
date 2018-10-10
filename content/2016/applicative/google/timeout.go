package google

import (
	"errors"
	"time"
)

func SearchTimeout(query string, timeout time.Duration) ([]Result, error) { // HL
	timer := time.After(timeout) // HL
	c := make(chan Result, 3)
	go func() { c <- Web(query) }()
	go func() { c <- Image(query) }()
	go func() { c <- Video(query) }()

	var results []Result
	for i := 0; i < 3; i++ {
		select { // HL
		case result := <-c: // HL
			results = append(results, result)
		case <-timer: // HL
			return results, errors.New("timed out")
		}
	}
	return results, nil
	// STOP OMIT
}
