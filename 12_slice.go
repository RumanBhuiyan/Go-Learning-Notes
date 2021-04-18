// slice allow programmer to create array which is resizable
package main

import "fmt"

func main(){
	// create an array which current size 3 but capacity 5
   var numbers = make([]int, 3, 5)

   numbers[0] = 5
   numbers[1] = 10

   fmt.Println("Length ",len(numbers))
   fmt.Println("capacity ",cap(numbers))

   // sub-array or slice 
   fmt.Println("first 2 elements ",numbers[0:2])

   // append function 
   numbers = append(numbers,3)

   fmt.Println(numbers)
}