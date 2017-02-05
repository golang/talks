package main

import (
	"fmt"
	"go/build"
	"log"
	"net/http"
	"path/filepath"
)

var cert, key string

func init() {
	pkg, err := build.Import("golang.org/x/talks/2017/state-of-go/stdlib/http2", ".", build.FindOnly)
	if err != nil {
		log.Fatal(err)
	}
	cert = filepath.Join(pkg.Dir, "cert.pem")
	key = filepath.Join(pkg.Dir, "key.pem")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/style.css", cssHandler)

	go func() {
		log.Fatal(http.ListenAndServeTLS("127.0.0.1:8081", cert, key, nil))
	}()
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	if p, ok := w.(http.Pusher); ok { // HL
		err := p.Push("/style.css", nil) // HL
		if err != nil {
			log.Printf("could not push: %v", err)
		}
	}

	fmt.Fprintln(w, html)
}

func cssHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, css)
}

const (
	html = `
<html>
<head>
	<link rel="stylesheet" href="/style.css">
	<title>HTTP2 push test</title>
</head>
<body>
	<h1>Hello</h1>
</body>
</html>
`
	css = `
h1 {
    color: red;
    text-align: center;
    text-shadow: green 0 0 40px;
    font-size: 10em;
}
`
)
