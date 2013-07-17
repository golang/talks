// +build OMIT

package main

func main() {
	var c1, c2, c3 chan int
	// START0 OMIT
	select {
	case v1 := <-c1:
		fmt.Printf("received %v from c1\n", v1)
	case v2 := <-c2:
		fmt.Printf("received %v from c2\n", v1)
	case c3 <- 23:
		fmt.Printf("sent %v to c3\n", 23)
	default:
		fmt.Printf("no one was ready to communicate\n")
	}
	// STOP0 OMIT
}
