// +build ignore,OMIT

package main

import "plugin"

func main() {
	p, err := plugin.Open("plugin_name.so")
	if err != nil {
		panic(err)
	}

	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}

	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}

	*v.(*int) = 7
	f.(func())() // prints "Hello, number 7"
}
