package main

import (
	"fmt"
)

// defining struct : type <Struct_name> struct {}
type Person struct {
	name string
	age int8
}

func main(){
	// process -01 : initializing struct instance with empty string and 0
	p1 := new(Person)
	fmt.Println("p1.name : ", p1.name,"p1.age : ", p1.age)

	//  process -02 : initializing struct instance with field names
	p2 := Person { name : "Ruman", age : 23}
	fmt.Println("p2.name : ", p2.name,"p2.age : ", p2.age)

	// process -03 : initializing struct instances without field names
	p3 := Person { "Ruman",  23}
	fmt.Println("p3.name : ", p3.name,"p3.age : ", p3.age)
}