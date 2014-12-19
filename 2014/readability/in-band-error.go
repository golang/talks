// +build OMIT

package sample // OMIT

import ( // OMIT
	"duration" // OMIT
	"time"     // OMIT
) // OMIT

// duration.Unterminated = -1 * time.Second

func (it Iterator) DurationAt() time.Duration {
	// some code
	switch durationUsec := m.GetDurationUsec(); durationUsec {
	case -1:
		return duration.Unterminated // HL
	case -2:
		return -2 // HL
	default:
		return time.Duration(durationUsec) * time.Microsecond // HL
	}
	return -3 // HL
}
