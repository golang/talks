// +build ignore,OMIT

package parser

// START OMIT
type ParsedFunc struct {
	text string
	eval func(float64) float64
}

func Parse(text string) (*ParsedFunc, error) {
	f, err := parse(text)
	if err != nil {
		return nil, err
	}
	return &ParsedFunc{text: text, eval: f}, nil
}

func (f *ParsedFunc) Eval(x float64) float64 { return f.eval(x) }
func (f *ParsedFunc) String() string         { return f.text }

// END OMIT
func parse(text string) (func(float64) float64, error) {
	return func(x float64) float64 { return x }, nil
}
