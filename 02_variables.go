// Go supports both statically typed and dynamic typed variables
package main

import "fmt"

func main(){
	// Statically Typed variable
	var name string = "go" 
	fmt.Println(name,"is an awesome language ")

	// dynamic typed variable
	var y = 5
	fmt.Printf("type of y  %T\n",y)
	fmt.Println("Value of y ",y)

	// multiple variable declarations in one line 
	var a,b,c = 5 ,4, 3
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
	fmt.Println("c = ",c)

	// declaring constant
	const PI float32 = 3.1416
	var radius float32 = 2.5
	fmt.Println("Area of the circle : ",PI*radius*radius)  

}