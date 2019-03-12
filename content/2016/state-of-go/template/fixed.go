// +build ignore,OMIT

package main

import (
	"html/template"
	"log"
	"os"
)

var tmpl = template.Must(template.New("tmpl").Parse(`
<ul>
{{range .}}<li>{{.}}</li>
{{end}}</ul>
`))

func main() {
	err := tmpl.Execute(os.Stdout, []string{"one", "two", "three"})
	if err != nil {
		log.Fatal(err)
	}
}
