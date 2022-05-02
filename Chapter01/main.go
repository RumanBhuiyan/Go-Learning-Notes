// Who are adopting Go are called gophers
// Go mascot logo is designed by Renee french

// package denotes  your go codes will  belong to where
package main

import (
	"fmt"
)

// for running any go file that file should have main() function so that go compiler
// can understand from where to start execution. Otherwise it will throw an error
func main(){
	fmt.Println("Hello world");
}

/*
	For getting any help on any function documentation 
	run the command   : go doc <package_name> <function_name>
	example			 : go doc fmt Println
*/