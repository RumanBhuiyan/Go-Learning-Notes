package main

import (
	"fmt"
)

func main() {

	// this is self invoking function
	func(){
	  fmt.Println("hello world")
	}()

	// holding function reference and using it
	greet := func(m string){
		fmt.Println(m)
	  }
	  
	  greet("Hey what's up!")
}
