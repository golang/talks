// +build OMIT

package main

import (
	"fmt"
	"reflect"
)

func makeSwap(fptr interface{}) {
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}
	fn := reflect.ValueOf(fptr).Elem()
	v := reflect.MakeFunc(fn.Type(), swap)
	fn.Set(v)
}

func main() {
	var fn func(int, int) (int, int)
	makeSwap(&fn)
	fmt.Println(fn(0, 1))
}
