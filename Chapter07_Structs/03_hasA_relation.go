package main

import (
	"fmt"
)

type Point struct{
	x float32
	y float32
}

/* At the time of creating Circle instances we dont have to create Point instance so 
	Circle isn't a Point. That means Circle has no 'is A' reletion with Point. But 
	Point.x = 0, Point.y =0 is initialized so that we can use it by c1.Point.x etc
	that means Now Circle has a 'has A' type relation with Point
*/
type Circle struct {
	radius float32
	Point
}

func main(){
	// c1 := Circle {
	// 	radius:  2.5,
	// }
	c1 := new(Circle)
	
	fmt.Println("c1.radius : ",c1.radius)
	fmt.Println("c1.Point.x : ",c1.Point.x)
	fmt.Println("c1.Point.y : ",c1.Point.y)
	
}