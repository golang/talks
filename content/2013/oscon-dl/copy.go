// +build ignore,OMIT

package main

import (
	"io"
)

func main() {
	var dst io.Writer
	var src io.Reader
	// START OMIT
	n, err := io.Copy(dst, src)
	// END OMIT
	_ = n
	_ = err
}
