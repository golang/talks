// +build OMIT

package main

import ( "fmt"; "net/http"; "time" )

func main() {
	urls := []string{"http://google.com/", "http://bing.com/"}
	start := time.Now()
	done := make(chan string)
	for _, u := range urls {
		go func(u string) {
			resp, err := http.Get(u)
			if err != nil {
				done <- u + " " + err.Error()
			} else {
				done <- u + " " + resp.Status
			}
		}(u)
	}
	for _ = range urls {
		fmt.Println(<-done, time.Since(start))
	}
}
