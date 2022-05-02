package main

import (
	"fmt"
)

func main(){
	numbers := [...]int { 1,2,3,3,4,5}
	fmt.Println("numbers : ",numbers)

	numbers_map := make(map[int]bool)
	for _, value := range numbers {
		numbers_map[value] = true
	}
	fmt.Println("numbers_map : ",numbers_map)

	var numbers2 []int
	for  key := range numbers_map {	
			numbers2 = append(numbers2, key)
	}
	fmt.Println("numbers2 : ",numbers2)
}