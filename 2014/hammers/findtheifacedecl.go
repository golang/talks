// +build ignore

package main

import (
	"go/ast"
	"go/build"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
)

func main() {
	fset, files := parsePackage("net/http")
	id := "Handler"

	for _, f := range files {
		for _, decl := range f.Decls {
			decl, ok := decl.(*ast.GenDecl)
			if !ok || decl.Tok != token.TYPE {
				continue
			}
			for _, spec := range decl.Specs {
				spec := spec.(*ast.TypeSpec)
				if spec.Name.Name == id {
					printer.Fprint(os.Stdout, fset, spec) // HL
				}
			}
		}
	}
}

func parsePackage(path string) (*token.FileSet, []*ast.File) {
	pkg, err := build.Import(path, "", 0)
	if err != nil {
		panic(err)
	}

	fset := token.NewFileSet()
	var files []*ast.File
	for _, file := range pkg.GoFiles {
		f, err := parser.ParseFile(fset, filepath.Join(pkg.Dir, file), nil, 0)
		if err != nil {
			continue
		}
		files = append(files, f)
	}
	return fset, files
}
