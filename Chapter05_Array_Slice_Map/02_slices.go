/*
	slices are growable sized array unlike array in Go
*/
package main

import (
	"fmt"
)

func main(){
	// creating slices 
	// process - 01
	var numbers []int32

	/* append doubles the current capacity of numbers if current size can't fit all items
	  and then keep previous items there and add new item there. Thus dynamic array works.
	  as vector is implemented on dynamic array so vector works like this procedure.
	  using make([]int , 0 ,10) , here 0 is length, 10 is capacity , thus we can tell go 
	  take 10 capacity and don't reallocate space until length exceeds this capacity 
	 Thus make() is used for pre-allocating space at the time of variable declaration 
	  */
	numbers = append(numbers,1)
	numbers = append(numbers,2)
	numbers = append(numbers,3)

	fmt.Println("numbers : ",numbers)
	fmt.Println("len of numbers : ",len(numbers))

	// accessing index and values simultaenously
	for index, value := range numbers {
		fmt.Println("index : ",index, "value : ",value)
	}

	// process - 02
	even_numbers := make([]int32,5,10) // create a slice of length 5 and capacity 10
	fmt.Println("Length of even_numbers : " ,len(even_numbers))
	fmt.Println("capacity of even_numbers : ", cap(even_numbers))

	for i:= 0 ; i < 5 ; i++{
		even_numbers[i] = int32(i*2)
	}
	fmt.Println("Even Numbers : ",even_numbers)
	fmt.Println("slices[1:]", even_numbers[1:])

	// copy function => copy(destination, source) overwrites the contents of destination by source
	copied_numbers := make([]int32,2)
	copy(copied_numbers, even_numbers[1:])
	fmt.Println(copied_numbers)

	/* using three index slices we can specify the capacity of slice like
	 slice_name[start_index : stop_index : capacity] */
}