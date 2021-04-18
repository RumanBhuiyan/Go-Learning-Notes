package main

import "fmt"

var PI float32 = 3.1416 
var age int = 23

func main(){
    var radius float32 ;
	fmt.Printf("Enter the radius: ")
	fmt.Scan(&radius)
	fmt.Println("Radius of the circle ",getArea(radius))

	// before creating same name variable of global scope
	fmt.Println("Global age: ",age)
	// after creating same name variable of global scope
	var age int =25
	fmt.Println("Local age: ",age)
}

func getArea(radius float32) float32 {
	// programmer didn't need to pass PI here for global scope
	return PI*radius*radius 
}