package main

import (
	"fmt"
	"math"
)

type Point struct{
	x float32
	y float32
}

// Here Circle is a Point, that means there is a 'is A' type relation between Point and Circle
// in 03_embed_types we will see how to implement a 'has A' type relation
type Circle struct {
	radius float32
	center Point
}

// defining method for Circle struct instances
func (c *Circle) area() float32 {
	return math.Pi * c.radius * c.radius 
}

func (c *Circle) set_radius(value float32) {
	c.radius = value
}

func main(){
	c1 := Circle {
		radius: 2.5,
		center: Point{
			x : 0.5,
			y : 0.5,
		},
	}

	fmt.Println("c1 area : ", c1.area())

	// c1.radius = 4
	c1.set_radius(float32(4))
	fmt.Println("c1.radius : ",c1.radius)
}