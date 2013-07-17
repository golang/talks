// +build OMIT

package main

import (
	"fmt"
	"reflect"
)

func sendSlice(slice interface{}) (channel interface{}) {
	sliceValue := reflect.ValueOf(slice)
	chanType := reflect.ChanOf(reflect.BothDir, sliceValue.Type().Elem())
	chanValue := reflect.MakeChan(chanType, 0)
	go func() {
		for i := 0; i < sliceValue.Len(); i++ {
			chanValue.Send(sliceValue.Index(i))
		}
		chanValue.Close()
	}()
	return chanValue.Interface()
}

func main() {
	ch := sendSlice([]int{1, 2, 3, 4, 5}).(chan int)
	for v := range ch {
		fmt.Println(v)
	}
}
