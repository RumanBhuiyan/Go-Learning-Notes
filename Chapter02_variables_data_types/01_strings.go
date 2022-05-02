package main

import (
	"fmt"
)

func main(){
	// declaring variables and assigning values
	// process - 01
	var name string = "Ruman"
	fmt.Println("name : "+name)

	// process - 02 (omit var keyword and type , let Go infer type on assignment)
	fullName := "Ruman Bhuiyan"
	fmt.Println("full name : "+fullName)

	// process - 03 (declare first then assign next)
	var country string
	country = "Bangladesh"
	fmt.Println("country : " + country)

	description1 := "Hey there! \n what's up" 
	description2 := ` Hey there ! 
		what's up again!
	`

	fmt.Println("description1 : " + description1)
	fmt.Println("description2 : " + description2)
}
/*
	N.B => string literals can be declared using "" or ``
	"" can't contain newline but accepts \n \t inside it
	`` accepts newline and \n \t also
*/