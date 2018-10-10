// +build OMIT

package P

import "fmt"

func before(s string) { fmt.Printf("%s\n", s) }
func after(s string)  { fmt.Println(s) }
