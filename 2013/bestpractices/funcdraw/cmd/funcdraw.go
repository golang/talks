// +build ignore,OMIT

package main

import (
	"flag"
	"image/png"
	"log"
	"os"
)

// IMPORT OMIT
import (
	"golang.org/x/talks/2013/bestpractices/funcdraw/drawer"
	"golang.org/x/talks/2013/bestpractices/funcdraw/parser"
)

// ENDIMPORT OMIT

var (
	width  = flag.Int("width", 300, "image width")
	height = flag.Int("height", 300, "image height")
	xmin   = flag.Float64("xmin", -10, "min value for x")
	xmax   = flag.Float64("xmax", 10, "max value for x")
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("missing expression to parse")
	}

	text := flag.Arg(0)
	// START OMIT
	// Parse the text into an executable function.
	f, err := parser.Parse(text)
	if err != nil {
		log.Fatalf("parse %q: %v", text, err)
	}

	// Create an image plotting the function.
	m := drawer.Draw(f, *width, *height, *xmin, *xmax)

	// Encode the image into the standard output.
	err = png.Encode(os.Stdout, m)
	if err != nil {
		log.Fatalf("encode image: %v", err)
	}
	// END OMIT
}
