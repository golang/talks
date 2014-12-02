// +build OMIT

package main

import (
	"fmt"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event"
	"golang.org/x/mobile/gl"
)

func main() {
	app.Run(app.Callbacks{
		Draw: func() {
			gl.ClearColor(0, 0, 1, 1) // blue
			gl.Clear(gl.COLOR_BUFFER_BIT)
		},
		Touch: func(e event.Touch) { fmt.Println(e) },
	})
}
