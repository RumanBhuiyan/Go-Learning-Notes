package main

import "fmt"

func main(){
	var numbers  = [5]int{5,10,15,20,25}
    var i int 
	for i = range numbers {
		fmt.Println(i+1,"element : ",numbers[i])
	}// output : 
	// 1 element :  5
	// 2 element :  10
	// 3 element :  15
	// 4 element :  20
	// 5 element :  25
}