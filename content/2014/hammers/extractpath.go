// +build ignore

package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"strconv"
)

func main() {
	src := `package hack; import "net/http"; var i http.Handler`

	fset := token.NewFileSet()
	f, _ := parser.ParseFile(fset, "", src, 0)

	raw := f.Imports[0].Path.Value
	path, _ := strconv.Unquote(raw)
	fmt.Println(raw, "\n", path)
}
