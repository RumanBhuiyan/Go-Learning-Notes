// array has fixed length size but slice length is resizable.

package main

import "fmt"

func main(){
	// Declaring slice 
	numbers := []int {1,2,3,4,5}
	countries := myArray{"USA","Canada","Bangladesh"}
	numberPhrases := giveEvenNumbers()

	// adding new item to slice
	countries = append(countries, "India");

	// accessing item 
	fmt.Println("Third country : ",countries[2])

	// looping over items
	for index,number := range numbers {
		fmt.Println(index,number)
	}

	// calling custom function
	countries.print()
	numberPhrases.print()

	// assigning multiple return statements
	firsthalf , secondhalf := provideHalf(numberPhrases,18)

	firsthalf.print()
	secondhalf.print()
}

