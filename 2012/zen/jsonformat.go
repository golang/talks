// +build OMIT

package main

import ( "encoding/json"; "fmt"; "io"; "os" )

func main() {
	d := json.NewDecoder(os.Stdin)
	var err error
	for err == nil {
		var v interface{}
		if err = d.Decode(&v); err != nil {
			break
		}
		var b []byte
		if b, err = json.MarshalIndent(v, "", "  "); err != nil {
			break
		}
		_, err = os.Stdout.Write(b)
	}
	if err != io.EOF {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
