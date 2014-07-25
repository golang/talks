// +build ignore

package main

import (
	"os"
	"text/template"
)

func main() {
	const stub = "func ({{.Recv}}) {{.Name}}" +
		"({{range .Params}}{{.Name}} {{.Type}}, {{end}})" +
		"({{range .Res}}{{.Name}} {{.Type}}, {{end}})" +
		"{\n}\n\n"
	tmpl := template.Must(template.New("test").Parse(stub))

	m := Method{
		Recv: "f *File",
		Func: Func{
			Name: "Close",
			Res:  []Param{{Type: "error"}},
		},
	}

	tmpl.Execute(os.Stdout, m)
}

// Method represents a method signature.
type Method struct {
	Recv string
	Func
}

// Func represents a function signature.
type Func struct {
	Name   string
	Params []Param
	Res    []Param
}

// Param represents a parameter in a function or method signature.
type Param struct {
	Name string
	Type string
}
