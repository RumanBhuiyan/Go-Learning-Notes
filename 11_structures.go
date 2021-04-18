package main

import "fmt"

type Info struct{
   name string
   age  int  
   occupation string
}

func main(){

	var first_person Info ;

	// assigning values
	first_person.name = "Ruman"
	first_person.age  = 24
	first_person.occupation = "student"

	// accessing items 
   fmt.Println("name: ",first_person.name)
   fmt.Println("age: ",first_person.age)
   fmt.Println("occupation: ",first_person.occupation)

}