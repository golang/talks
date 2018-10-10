// +build ignore

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	src := `
package http
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
`
	f, _ := parser.ParseFile(token.NewFileSet(), "", src, 0)
	typ := f.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Type.(*ast.InterfaceType)
	fndecl := typ.Methods.List[0].Type.(*ast.FuncType)
	// fndecl: (ResponseWriter, *Request)

	ast.Inspect(fndecl, func(n ast.Node) bool { // HL
		if ident, ok := n.(*ast.Ident); ok {
			fmt.Println(ident.Name)
		}
		return true
	})
}

// end main // OMIT
