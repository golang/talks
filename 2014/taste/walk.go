// +build OMIT

package main

import (
	"fmt"
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

func _() {
	// example START OMIT
	n := 0
	println := func(s string) bool {
		fmt.Println(n, s)
		n++
		return n < 10
	}
	walkStdLib(println)
	// example END OMIT
}

func main() {
	// main START OMIT
	n := 0
	walkStdLib(func(s string) bool {
		fmt.Println(n, s)
		n++
		return n < 10
	})
	// main END OMIT
}
