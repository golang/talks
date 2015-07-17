package main

var global *int

func f(i *int) { global = i }

func main() {
	a := new(int)
	f(a)
}
