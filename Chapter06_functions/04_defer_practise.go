package main

import (
	"fmt"
	"os"
)

func main(){
	// this will call first f1() then f2()
	// f1()
	// f2()

	// defer will call f1() at last of all statements of main()
	// suppose you opened a file and forget to close in end in such case call defer file.Close()
	// this will close your file at end of function
	defer f1()
	f2()

	file,_ := os.Open("./test.txt")
	defer file.Close()
}

func f1(){
	fmt.Println("func f1 called!")
}

func f2(){
	fmt.Println("func f2 called!")
}