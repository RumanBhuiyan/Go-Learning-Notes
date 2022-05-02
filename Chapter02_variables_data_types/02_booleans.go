package main

import (
	"fmt"
)

func main(){
	var track1 bool = true
	var track2 bool = false

	fmt.Println(track1 && track2)
	fmt.Println(track1 || track2)
	fmt.Println(!track1)
}
/*
	Each hexadecimal digits requires 4 bits to be represented as hexadecimal 
	digits ranges from 0-15. 4 bits are called one nibble
*/