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
	"sort"
	"strings"
	"time"
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

type histogram map[string]int

func (h histogram) add(filename string) {
	f, err := parser.ParseFile(token.NewFileSet(), filename, nil, 0)
	if err != nil {
		panic(err)
	}

	ast.Inspect(f, func(n ast.Node) bool {
		if n, ok := n.(ast.Stmt); ok {
			h[fmt.Sprintf("%T", n)]++
		}
		return true
	})
}

// print START OMIT
func (h histogram) print() {
	var list []entry
	var total int
	for key, count := range h {
		list = append(list, entry{key, count})
		total += count
	}
	sort.Sort(byCount(list))

	percent := 100 / float64(total)
	for i, e := range list {
		fmt.Printf("%4d.  %5.2f%%  %5d  %s\n", i, float64(e.count)*percent, e.count, e.key)
	}
}

// print END OMIT

// byCount START OMIT
type entry struct {
	key   string
	count int
}

type byCount []entry

func (s byCount) Len() int      { return len(s) }
func (s byCount) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s byCount) Less(i, j int) bool {
	x, y := s[i], s[j]
	if x.count != y.count {
		return x.count > y.count // want larger count first
	}
	return x.key < y.key
}

// byCount END OMIT

// main START OMIT
func main() {
	start := time.Now()
	h := make(histogram)
	walkStdLib(func(filename string) bool {
		h.add(filename)
		return true
	})

	h.print()
	fmt.Println(time.Since(start))
}

// main END OMIT
