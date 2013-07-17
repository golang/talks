// +build OMIT

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string
type Search func(query string) Result

var (
	Web = fakeSearch("web")
	Image = fakeSearch("image")
	Video = fakeSearch("video")
)


func Google(query string) (results []Result) {
	c := make(chan Result)
	go func() { c <- Web(query) } ()
	go func() { c <- Image(query) } ()
	go func() { c <- Video(query) } ()

	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

func fakeSearch(kind string) Search {
        return func(query string) Result {
	          time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	          return Result(fmt.Sprintf("%s result for %q\n", kind, query))
        }
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Google("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}


