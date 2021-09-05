package main

import "fmt"

// here i am creating my own type called 'myArray'
// which is basically a slice of strings
type myArray []string

// i am adding print method with my type so that i can access it like
// country.print()
func (c myArray) print(){
	for index, country := range c{
		fmt.Println(index,country);
	}
}

func giveEvenNumbers() myArray{
	var keep myArray
	
	numbers :=[]string{"tweenty","thirty","forty","fifty"}
	phrases :=[]string{"one","two","three","four","five","six","seven","eight","nine"}

	// i am not going to use the index here so using _ to ignore it
	for _,number := range numbers{
			for _,phrase := range phrases{
				keep = append(keep,number+" "+phrase)
			}
	}

	return keep;
}

// returning multiple statements
func provideHalf(a myArray, size int) (myArray,myArray){
	return a[:size], a[size:]
}