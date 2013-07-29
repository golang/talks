// +build ignore,OMIT

package drawer

// START OMIT
import "image"

// Function represent a drawable mathematical function.
type Function interface {
	Eval(float64) float64
}

// Draw draws an image showing a rendering of the passed Function.
func Draw(f Function) image.Image {
	// END OMIT
	return nil
}
