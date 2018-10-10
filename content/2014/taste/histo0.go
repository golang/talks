// +build OMIT

package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

func walk(dir string, f func(string) bool) bool {
	fis, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	// parse all *.go files in directory;
	// traverse subdirectories, but don't walk into testdata
	for _, fi := range fis {
		path := filepath.Join(dir, fi.Name())
		if fi.IsDir() {
			if fi.Name() != "testdata" {
				if !walk(path, f) {
					return false
				}
			}
		} else if strings.HasSuffix(fi.Name(), ".go") && !strings.HasPrefix(fi.Name(), ".") {
			if !f(path) {
				return false
			}
		}
	}
	return true
}

func walkStdLib(f func(filename string) bool) {
	walk(filepath.Join(runtime.GOROOT(), "src"), f)
}

// histogram START OMIT
type histogram map[string]int

// histogram END OMIT

// add START OMIT
func (h histogram) add(filename string) {
	f, err := parser.ParseFile(token.NewFileSet(), filename, nil, 0)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		if n, ok := n.(ast.Stmt); ok { // type test: is n an ast.Stmt?
			h[fmt.Sprintf("%T", n)]++
		}
		return true
	})
}

// add END OMIT

// print START OMIT
func (h histogram) print() {
	// determine total number of statements
	total := 0
	for _, count := range h {
		total += count
	}

	// print map entries
	i := 0
	percent := 100 / float64(total)
	for key, count := range h {
		fmt.Printf("%4d.  %5.2f%%  %5d  %s\n", i, float64(count)*percent, count, key)
		i++
	}
}

// print END OMIT

// main START OMIT
func main() {
	// body START OMIT
	h := make(histogram)
	walkStdLib(func(filename string) bool {
		h.add(filename) // does all the hard work
		return true
	})
	// body END OMIT
	h.print()
}

// main END OMIT
