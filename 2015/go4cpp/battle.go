// +build OMIT

package main

import (
	"fmt"
	"net/http"
)

var battle = make(chan string)

func handler(w http.ResponseWriter, q *http.Request) {
	select {
	case battle <- q.FormValue("usr"):
		fmt.Fprintf(w, "You won!")
	case won := <-battle:
		fmt.Fprintf(w, "You lost, %v is better than you", won)
	}
}

func main() {
	http.HandleFunc("/fight", handler)
	http.ListenAndServe("localhost:8080", nil)
}
