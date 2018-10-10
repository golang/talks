// +build ignore

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	src := `package hack; import "net/http"; var i http.Handler`
	f, _ := parser.ParseFile(token.NewFileSet(), "", src, 0)

	decl := f.Decls[1].(*ast.GenDecl)      // var i http.Handler
	spec := decl.Specs[0].(*ast.ValueSpec) // i http.Handler
	sel := spec.Type.(*ast.SelectorExpr)   // http.Handler
	id := sel.Sel.Name                     // Handler
	fmt.Println(id)
}
