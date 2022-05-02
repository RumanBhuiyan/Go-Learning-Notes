package main

import (
	"fmt"
)

func main(){
	fmt.Println(" 3 / 2 = ",get_quotient(float32(3),float32(2)))
	fmt.Println(" 3 / 0 = ",get_quotient(float32(3),float32(0)))
}

func get_quotient(dividend, divisor float32) float32 {
	if divisor != 0{
		result := dividend / divisor
		return result
	}
	/* this func will be executed at last so if panic(message) is called then
	 that message can be received by recover(). As a runtime recovery we can handle
	 this type of panic by writing our own function like func below */
	defer func(){
		message := recover()
		fmt.Println("Error message : ",message)
	}()
	panic("Can't divide by zero")
}