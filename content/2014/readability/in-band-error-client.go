// +build OMIT

package client // OMIT

func proc(it Iterator) (ret time.Duration) {
	d := it.DurationAt()
	if d == duration.Unterminated { // HL
		ret = -1
	} else {
		ret = d
	}
	// some code
}
