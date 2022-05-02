package main

import (
	"fmt"
	"os"
)

func main(){
	args := os.Args
	fmt.Println("command line arguments : ",args)

	for _, value := range args[1:] {
		fmt.Println("You provided ",value, " as input")
	}
}