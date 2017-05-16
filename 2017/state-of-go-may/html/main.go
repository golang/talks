package main

import (
	"html/template"
	"log"
	"os"
)

type Foo struct{ Bar string }

func main() {
	tmpl, err := template.New("home").Parse(`
		<a title={{.Bar | html}}>
	`)
	if err != nil {
		log.Fatalf("could not parse: %v", err)
	}

	foo := Foo{"haha onclick=evil()"}
	if err := tmpl.Execute(os.Stdout, foo); err != nil {
		log.Fatalf("could not execute: %v", err)
	}
}
