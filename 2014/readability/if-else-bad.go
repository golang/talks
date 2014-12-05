// +build OMIT

package sample // OMIT

func sample() { // OMIT
	if _, ok := f.dirs[dir]; !ok { // HL
		f.dirs[dir] = new(feedDir) // HL
	} else {
		f.addErr(fmt.Errorf("..."))
		return
	}
	// some code
} // OMIT
