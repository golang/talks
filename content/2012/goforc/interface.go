// +build OMIT

package main

import "fmt"

type Point struct{ x, y int }

func (p Point) String() string { return fmt.Sprintf("Point{%d, %d}", p.x, p.y) }

type Celsius float32
type Fahrenheit float32

func (t Celsius) String() string           { return fmt.Sprintf("%g°C", t) }
func (t Fahrenheit) String() string        { return fmt.Sprintf("%g°F", t) }
func (t Celsius) ToFahrenheit() Fahrenheit { return Fahrenheit(t*9/5 + 32) }

func main() {
	// START OMIT
	type Stringer interface {
		String() string
	}

	var v Stringer // HL
	var corner = Point{1, 1}
	var boiling = Celsius(100)

	v = corner
	fmt.Println(v.String()) // dynamic dispatch
	fmt.Println(v)

	v = boiling.ToFahrenheit()
	fmt.Println(v.String()) // dynamic dispatch
	fmt.Println(v)
	// STOP OMIT
}
