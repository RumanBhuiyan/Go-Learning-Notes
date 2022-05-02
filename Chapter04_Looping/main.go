package main

import (
	"fmt"
)

func main(){
	// basic for loop
	for i:= 1; i<=5; i++ {
		fmt.Println("i : ",i)
	}

	// while style for loop
	i:= 6
	for i<= 10 {
		fmt.Println("i : ",i)
		i += 1
	}

	// creating an infinite for loop
	for {
		fmt.Println("Hello!")
	}
}
/*
	N.B => do-while, while, untill, forEach loops are unavailable in go
	You will have to do all looping stuffs using for loop in Go 
*/