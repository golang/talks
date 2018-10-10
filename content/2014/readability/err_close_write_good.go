// +build OMIT

package sample // OMIT

func run() (err error) {
	in, err := os.Open(*input)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(*output)
	if err != nil {
		return err
	}
	defer func() { // HL
		if cerr := out.Close(); err == nil { // HL
			err = cerr // HL
		} // HL
	}() // HL
	// some code
}
