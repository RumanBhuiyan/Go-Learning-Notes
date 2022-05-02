/* 
	map in Go is equivalent of dictionary in Python and hashmap, associative arrays in other
	programming language data structures
*/
package main

import (
	"fmt"
)

func main(){
	// make() preallocates space so that we can access and alter values by keys
	chemical_name := make(map[string] string)

	chemical_name["H"] = "Hydrozend"
	chemical_name["He"] = "Helium"
	chemical_name["Li"] = "Lithyum"
	chemical_name["Be"] = "Berilyum"
	chemical_name["B"] = "Boron"
	chemical_name["C"] = "Carbon"
	chemical_name["N"] = "Nitrozen"
	
	fmt.Println(chemical_name)

	for key,value := range chemical_name {
		fmt.Println("Key : ",key, "\t| Value : ",value)
	}

	// accessing key from a map in a safe way 
	if value,ok := chemical_name["O"]; ok{
		fmt.Println("Key found value : ",value)
	}else {
		fmt.Println("key not found")
	}

	// deleting a key from map
	delete(chemical_name,"N")
	
	fmt.Println(chemical_name)

	// map is passed by reference in assignment that means any change in anyone is reflected in both
	numbers  := make(map[string] int )
	
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	
	fmt.Println(numbers)
	
	n := numbers
	
	n["Ruman"] = 4
	
	fmt.Println(n)
	fmt.Println(numbers)
}