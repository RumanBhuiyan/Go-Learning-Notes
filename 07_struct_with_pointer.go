package main

import "fmt"

type person struct {
	fname string
}

// Receiver function
func (p person) updateName(name string) {
	p.fname = name
}

// Receiver function with pointer 
func (personPointer *person) updateName2(name string){
    (*personPointer).fname = name
}

func main() {

	ruman := person{fname: "Ruman"}
	alex := person {fname: "Alex"}
	//fmt.Println(ruman.fname)

	// Both of the output will be Ruman and Ruman
	// because a copy of current struct is passed to receiver function
	// as actual reference isn't passed thats why actual value isn't changed
	// to solve this problem use pointer.
	ruman.updateName("Bhuiyan")
	//fmt.Println(ruman.fname)

	rumanPointer := &ruman
	rumanPointer.updateName2("Bhuiyan")
	fmt.Println(ruman.fname)

		//Or
	alex.updateName2("Balex")
	fmt.Println(alex.fname)
}
