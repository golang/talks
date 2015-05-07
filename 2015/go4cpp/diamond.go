// +build OMIT

package main

import "fmt"

type Engine struct{}

func (e Engine) Start() { fmt.Println("Engine started") }
func (e Engine) Stop()  { fmt.Println("Engine stopped") }

type Radio struct{}

func (r Radio) Start() { fmt.Println("Radio started") }
func (r Radio) Stop()  { fmt.Println("Radio stopped") }

type Car struct {
	Engine
	Radio
}

func main() {
	var c Car
	c.Radio.Start()  // HL
	c.Engine.Start() // HL
}
