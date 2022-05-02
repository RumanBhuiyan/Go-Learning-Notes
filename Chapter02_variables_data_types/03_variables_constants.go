package main

import (
	"fmt"
)

// defining constants
const PI float32 = 3.1416

func main(){

	// defining multiple variables together
	var (
		a = 10
		b = 20
		c = 30
	)

	fmt.Println("a : " , a)
	fmt.Println("a : " , b)
	fmt.Println("a : " , c)
	
	whatever()
	fmt.Println("value of PI : ", PI) 
}
func whatever(){
	fmt.Println("value of PI : ", PI) 
}