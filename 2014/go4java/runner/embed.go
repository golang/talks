package runner

// RunCounter2 is completely equivalent to RunCounter,
// but uses struct embedding to avoid the boilerplate of redeclaring
// the Name method.
type RunCounter2 struct {
	Runner // HL
	count  int
}

func NewRunCounter2(name string) *RunCounter2 {
	return &RunCounter2{Runner{name}, 0}
}

func (r *RunCounter2) Run(t Task) {
	r.count++
	r.Runner.Run(t) // HL
}

func (r *RunCounter2) RunAll(ts []Task) {
	r.count += len(ts)
	r.Runner.RunAll(ts) // HL
}

func (r *RunCounter2) Count() int { return r.count }
