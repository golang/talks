// +build ignore,OMIT

package main

import "net/http"

func main() {
	http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeContent(w, r, "foo.txt", time.Now(),
			strings.NewReader("I am some content.\n"))
	}))
}
