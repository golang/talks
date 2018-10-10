// Package runner provides a Runner type that is used to define both RunCounter
// and EmbeddedRunCounter to show examples of how to use composition in Go.
package runner

import "fmt"

// A Task is a simple task that prints a message when run.
type Task struct{ Msg string }

func (t Task) Run() {
	fmt.Println("running", t.Msg)
}

// END_TASK OMIT

// A Runner provides a way of running tasks.
type Runner struct{ name string }

func (r *Runner) Name() string { return r.name }

func (r *Runner) Run(t Task) {
	t.Run()
}

func (r *Runner) RunAll(ts []Task) {
	for _, t := range ts {
		r.Run(t)
	}
}

// END_RUNNER OMIT

// A RunCounter is a Runner that keeps a counter of the run tasks.
type RunCounter struct {
	runner Runner // HL
	count  int
}

func NewRunCounter(name string) *RunCounter {
	return &RunCounter{runner: Runner{name}}
}

func (r *RunCounter) Run(t Task) {
	r.count++
	r.runner.Run(t) // HL
}

func (r *RunCounter) RunAll(ts []Task) {
	r.count += len(ts)
	r.runner.RunAll(ts) // HL
}

func (r *RunCounter) Count() int { return r.count }

func (r *RunCounter) Name() string { return r.runner.Name() }
