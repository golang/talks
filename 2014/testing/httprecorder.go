// +build OMIT

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
)

func main() {
	// START OMIT
	handler := func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "something failed", http.StatusInternalServerError)
	}

	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		log.Fatal(err)
	}

	w := httptest.NewRecorder()
	handler(w, req)

	fmt.Printf("%d - %s", w.Code, w.Body.String())
	// STOP OMIT
}
