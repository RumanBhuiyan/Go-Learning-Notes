package main

import (
	"fmt"
)

func main(){

	// add_one := increment(5)
	add_one := increment(0)

	fmt.Println(add_one())
	fmt.Println(add_one())
	fmt.Println(add_one())
}

func increment(initial int32) func()int32 {
	current := initial

	fn := func() int32{
		current += 1
		return current 
	}
	return fn 
}