// +build ignore,OMIT

package drawer

import (
	"math"
	"testing"
)

type TestFunc func(float64) float64

func (f TestFunc) Eval(x float64) float64 { return f(x) }

var (
	ident = TestFunc(func(x float64) float64 { return x })
	sin   = TestFunc(math.Sin)
)

func TestDraw_Ident(t *testing.T) {
	m := Draw(ident)
	// Verify obtained image.
	// END OMIT
	t.Error(m.ColorModel())
}
