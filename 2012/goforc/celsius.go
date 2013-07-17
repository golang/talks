// +build OMIT

package main

import "fmt"

type Celsius float32
type Fahrenheit float32

func (t Celsius) String() string           { return fmt.Sprintf("%g°C", t) }
func (t Fahrenheit) String() string        { return fmt.Sprintf("%g°F", t) }
func (t Celsius) ToFahrenheit() Fahrenheit { return Fahrenheit(t*9/5 + 32) }

func main() {
	var t Celsius = 21
	fmt.Println(t.String())
	fmt.Println(t)
	fmt.Println(t.ToFahrenheit())
}
