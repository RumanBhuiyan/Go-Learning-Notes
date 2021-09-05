package main

import "fmt"

// if any structure has a function named printArea() then
// that will be Shape type and here i am declaring this interface
// because other members should use it
type Shape interface {
	printArea()
}

type Rectangle struct {
	length float64
	width  float64
}

type Square struct {
	armlength float64
}

// let make Rectangle Shape type so that it can call rectangle.printArea()
func (r Rectangle) printArea() {
	fmt.Println(r.length * r.width)
}

// let make Square Shape type so that it can square.printArea()
func (s Square) printArea() {
	fmt.Println(s.armlength * s.armlength)
}

// As Rectangle and Square both are Shape type and both of them have
// their own implementation of printArea() function , so now if we
// call printArea() then basis on type compile will decide which function
// it should call
func printDistinctArea(s Shape) {
	s.printArea()
}

func main() {

	rectangle := Rectangle{length: 2.5, width: 3.2}
	square := Square{armlength: 5}

	// now i can call same name function being concerned about their type
	rectangle.printArea()
	square.printArea()
		// or i can tell compiler to decide which function it should call
	printDistinctArea(rectangle)
	printDistinctArea(square)

}
