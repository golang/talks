package main

import "golang.org/x/net/context" // HL

func main() {
	ctx := context.Background()
	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	// doing something
}
