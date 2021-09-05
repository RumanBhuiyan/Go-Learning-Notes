package main

import "fmt"

func main(){

	// declaring variables 
	var fname string = "Ruman"
	// use := only at initialization time
	// lname := "Bhuiyan" is similar as var lname string ; lname = "Bhuiyan"
	lname := "Bhuiyan"
	fmt.Println("Hello ",fname," ",lname)

	// use = in re-assignmet time
	lname = "Ahmed"
	fmt.Println("Hello ",fname," ",lname)

	// printing value and  type 
	fmt.Printf("value: %v Type : %T \n",fname,fname)

	// storing value and type
	fnameValue := fmt.Sprintf("%v",fname)
	fnameType := fmt.Sprintf("%T",fname)

	fmt.Printf("Fname Value : %v Type : %T \n",fnameValue,fnameType)

}