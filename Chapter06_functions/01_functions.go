package main

import (
	"fmt"
)

func main(){
	numbers := [5]int32{1,2,3,4,5}
	fmt.Println("Average : ", calculate_average(numbers[:]))

	x,y := get_point();
	fmt.Println("x : ",x, "y : ",y)

	fmt.Println("summation of ",numbers," is ",get_summation(numbers[:]))
}

// functions name always should be started with capital letter otherwise package wont find that function

func calculate_average(nums []int32) float32 {
	total := int32(0)
	for _,value := range nums {
		total += value
	}
	return float32(total) / float32(len(nums))
}

// function returning multiple values
func get_point() (float32, float32) {
	return 0.5,0.2
}

// function that returns named value
func get_summation(nums []int32) (sum int32){
	sum = int32(0)
	for _,value := range nums{
		sum += value
	}
	return
}