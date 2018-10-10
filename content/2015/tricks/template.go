// +build ignore

package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

func main() {
	tmpl := template.Must(template.New("").Parse(strings.TrimSpace(`
Dear {{.Title}} {{.Lastname}},

Congratulations on reaching Level {{.Rank}}!
I'm sure your parents would say "Great job, {{.Firstname}}!"

Sincerely,
Rear Admiral Gopher
	`)))
	// BEGIN OMIT
	data := struct {
		Title               string
		Firstname, Lastname string
		Rank                int
	}{
		"Dr", "Carl", "Sagan", 7,
	}
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
	// END OMIT
}
