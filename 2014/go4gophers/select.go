package main

import (
	"net/http"
	"text/template"
	"time"
)

type Result struct{}

func WebSearch(q string) *Result   { return nil }
func ImageSearch(q string) *Result { return nil }

var searchTemplate *template.Template

func searchHandler(w http.ResponseWriter, r *http.Request) {
	var (
		results []*Result
		ch      = make(chan *Result)
		timeout = time.After(20 * time.Millisecond)
	)
	go func() { ch <- WebSearch(r) }()
	go func() { ch <- ImageSearch(r) }()
loop:
	for i := 0; i < 2; i++ {
		select { // HL
		case r := <-ch: // HL
			results = append(results, r)
		case <-timeout: // HL
			break loop
		}
	}
	searchTemplate.Execute(w, results)
}
