package main

import (
	"fmt"
)

func main(){
	// declaring arrays
	var numbers [5]int32
	fmt.Println(numbers)

	// accessing values and changing its values
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30
	numbers[3] = 40
	numbers[4] = 50
	fmt.Println(numbers)

	// use range as an alternative of enumerate() of python and Rust
	for index, value := range numbers {
		fmt.Println("index : ",index, "value : ",value)
	}

	summation := int32(0)
	for _,value := range numbers {
		summation += value
	}

	// N.B => so simple Type casting is  type(data)
	fmt.Println("Average : ",float32(summation) / float32(len(numbers)))

	// another way of creating aray
	even_numbers := [5]int32{
		2,
		4,
		6,
	}
	fmt.Println(even_numbers)
	fmt.Println(len(even_numbers))

	// declaring arrays let go infer its length
	planets := [...]string{
		"mercury",
		"venus",
		"earth",
	}
	// N.B => planets2 gets a copy of planets not reference
	planets2 := planets
	planets[2] = "mars"

	fmt.Println("planets : ",planets)
	fmt.Println("planets2 : ",planets2)
}