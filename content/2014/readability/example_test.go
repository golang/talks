// +build OMIT

package binary // OMIT

func ExampleWrite() {
	var buf bytes.Buffer
	var pi float64 = math.Pi
	err := binary.Write(&buf, binary.LittleEndian, pi)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x", buf.Bytes())
	// Output: 18 2d 44 54 fb 21 09 40
}
