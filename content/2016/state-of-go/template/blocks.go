// +build ignore,OMIT

package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

const tmplText = `
{{define "presentation"}}
	Authors:
	{{block "list" .Authors}} // HL
	<ul>
	{{- range .}}
		<li>{{.}}</li>
	{{- end}}
	</ul>
	{{end}} // HL

	Topics:
	{{template "list" .Topics}} // HL
{{end}}
` // OMIT

type Presentation struct {
	Authors []string
	Topics  []string
}

func main() {
	p := Presentation{
		Authors: []string{"one", "two", "three"},
		Topics:  []string{"go", "templates"},
	}

	tmpl := template.Must(template.New("presentation").Parse(tmplText))

	tmpl = tmpl.Funcs(template.FuncMap{"join": strings.Join})                       // HL
	tmpl = template.Must(tmpl.Parse(`{{define "list"}} {{join . " | "}} {{ end}}`)) // HL

	err := tmpl.Execute(os.Stdout, p)
	if err != nil {
		log.Fatal(err)
	}
}
