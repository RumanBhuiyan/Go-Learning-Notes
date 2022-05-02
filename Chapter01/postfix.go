// prefix isn't supported in go, but postfix is supported in Go
package main

import "fmt"

func main(){
	a := 5
	a++
	fmt.Println("a : ",a)
}