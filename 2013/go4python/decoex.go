// +build OMIT

package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

type errorHandler func(http.ResponseWriter, *http.Request) error

func handleError(f errorHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Oops!", http.StatusInternalServerError)
		}
	}
}

func handler(w http.ResponseWriter, r *http.Request) error {
	name := r.FormValue("name")
	if name == "" {
		return fmt.Errorf("empty name")
	}
	fmt.Fprintln(w, "Hi,", name)
	return nil
}

// resp implements http.ResponseWriter writing
type dummyResp struct {
	io.Writer
	h int
}

func newDummyResp() http.ResponseWriter {
	return &dummyResp{Writer: &bytes.Buffer{}}
}

func (w *dummyResp) Header() http.Header { return make(http.Header) }
func (w *dummyResp) WriteHeader(h int)   { w.h = h }
func (w *dummyResp) String() string      { return fmt.Sprintf("[%v] %q", w.h, w.Writer) }

func main() {
	http.HandleFunc("/hi", handleError(handler))

	// ListenAndServe is not allowed on the playground.
	// http.ListenAndServe(":8080", nil)

	// In the playground we call the handler manually with dummy requests.

	// Fake request without 'name' parameter.
	r := &http.Request{}
	w := newDummyResp()
	handleError(handler)(w, r)
	fmt.Println("resp a:", w)

	// Fake request with 'name' parameter 'john'.
	r.Form["name"] = []string{"john"}
	w = newDummyResp()
	handleError(handler)(w, r)
	fmt.Println("resp b:", w)

}
