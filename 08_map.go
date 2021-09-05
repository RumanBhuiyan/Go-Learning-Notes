package main

import "fmt"

func main() {

	// declaration process : 01
	var names map[int]string
	names = make(map[int]string)

	names[0] = "One"
	names[1] = "Two"
	names[2] = "Three"

	fmt.Println(names)

	// declaration process : 02
	colors := make(map[string]string)

	colors["red"] = "good"
	colors["green"] = "better"
	colors["blue"] = "best"

	printItems(colors)

	//deleting any key
	delete(colors,"red")
	printItems(colors)
}

func printItems (k map[string]string){
	// iterating over items 
	for key,value := range k{
		fmt.Println(key,value)
	}
}
