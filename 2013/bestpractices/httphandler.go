// +build ignore,OMIT

package bestpractices

import (
	"fmt"
	"log"
	"net/http"
)

func doThis() error { return nil }
func doThat() error { return nil }

// HANDLER1 OMIT
func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	err := doThis()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("handling %q: %v", r.RequestURI, err)
		return
	}

	err = doThat()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Printf("handling %q: %v", r.RequestURI, err)
		return
	}
}

// HANDLER2 OMIT
func init() {
	http.HandleFunc("/", errorHandler(betterHandler))
}

func errorHandler(f func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			log.Printf("handling %q: %v", r.RequestURI, err)
		}
	}
}

func betterHandler(w http.ResponseWriter, r *http.Request) error {
	if err := doThis(); err != nil {
		return fmt.Errorf("doing this: %v", err)
	}

	if err := doThat(); err != nil {
		return fmt.Errorf("doing that: %v", err)
	}
	return nil
}

// END OMIT
