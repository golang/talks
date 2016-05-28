package google

import (
	"fmt"
	"math/rand"
	"time"
)

// START2 OMIT
var (
	Web   = FakeSearch("web", "The Go Programming Language", "http://golang.org")
	Image = FakeSearch("image", "The Go gopher", "https://blog.golang.org/gopher/gopher.png")
	Video = FakeSearch("video", "Concurrency is not Parallelism", "https://www.youtube.com/watch?v=cN_DpYBzKso")
)

type SearchFunc func(query string) Result // HL

func FakeSearch(kind, title, url string) SearchFunc {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond) // HL
		return Result{
			Title: fmt.Sprintf("%s(%q): %s", kind, query, title),
			URL:   url,
		}
	}
}

// STOP2 OMIT

// String returns the result's title, followed by a newline.
func (r Result) String() string { return r.Title + "\n" }

var (
	Web1   = FakeSearch("web1", "The Go Programming Language", "http://golang.org")
	Web2   = FakeSearch("web2", "The Go Programming Language", "http://golang.org")
	Image1 = FakeSearch("image1", "The Go gopher", "https://blog.golang.org/gopher/gopher.png")
	Image2 = FakeSearch("image2", "The Go gopher", "https://blog.golang.org/gopher/gopher.png")
	Video1 = FakeSearch("video1", "Concurrency is not Parallelism",
		"https://www.youtube.com/watch?v=cN_DpYBzKso")
	Video2 = FakeSearch("video2", "Concurrency is not Parallelism",
		"https://www.youtube.com/watch?v=cN_DpYBzKso")
)
