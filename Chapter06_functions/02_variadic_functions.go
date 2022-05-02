package main

import (
	"fmt"
)

func main(){
	numbers := []int32 { 10, 20, 30, 40, 50}

	fmt.Println("summation of 1,2,3 : ", get_sum(1,2,3))
	fmt.Println("summation of 1,2,3,4,5 : ", get_sum(1,2,3,4,5))
	fmt.Println("summation of ",numbers," : ", get_sum(numbers...))
}

// variadic functions can take variable  arguments and process them
func get_sum(nums ...int32) int32 {
	summation := int32(0)
	for _,value := range nums {
		summation += value
	}
	return summation
}