package main

import "fmt"
import myFuncs "./lib"

func main(){
	fmt.Println("2 + 3 = ",myFuncs.GetSum(2,3))
}