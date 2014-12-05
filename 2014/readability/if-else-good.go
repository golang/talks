// +build OMIT

package sample // OMIT

func sample() { // OMIT

	if _, found := f.dirs[dir]; found { // HL
		f.addErr(fmt.Errorf("..."))
		return
	}
	f.dirs[dir] = new(feedDir) // HL
	// some code
} // OMIT
