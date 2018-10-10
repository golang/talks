// +build OMIT

package sample // OMIT

func run() error {
	in, err := os.Open(*input)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(*output)
	if err != nil {
		return err
	}
	defer out.Close() // HL
	// some code
}
