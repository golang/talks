// +build OMIT

package sample // OMIT

import ( // OMIT
	"errors" // OMIT
	"time"   // OMIT
) // OMIT

var (
	ErrDurationUnterminated = errors.new("duration: unterminated")
	ErrNoDuration           = errors.New("duration: not found")
	ErrNoIteration          = errors.New("duration: not interation")
)

func (it Iterator) DurationAt() (time.Duration, error) { // HL
	// some code
	switch durationUsec := m.GetDurationUsec(); durationUsec {
	case -1:
		return 0, ErrDurationUnterminated // HL
	case -2:
		return 0, ErrNoDuration // HL
	default:
		return time.Duation(durationUsec) * time.Microsecond, nil // HL
	}
	return 0, ErrNoIteration // HL
}
