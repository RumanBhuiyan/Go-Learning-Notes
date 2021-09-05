
package main

import "fmt"

type person struct{
	fname string
	lname string
	age int
	contact contactInfo
}

type contactInfo struct{
	mobile string
	email string
}

func main(){
	var alex, Ruman person

	// initailizing struct 
	 Ruman = person{fname:"Ruman",lname: "Bhuiyan",age:24}
	 alex = person {fname: "Alex", lname: "Anderson",age: 32}

	// accessing struct members
	fmt.Println(Ruman) // prints only values not key names
	fmt.Printf("%+v",Ruman)

	// updating values
	alex.age = 27

	fmt.Printf("%+v",alex)

	// lets initailize nested struct 
	elon := person{
		fname: "Elon",
		lname: "Musk",
		age: 32,
		contact: contactInfo{
			mobile: "12356",
			email: "ëlon@gmail.com",
		},
	}

	fmt.Printf("%+v",elon)

}