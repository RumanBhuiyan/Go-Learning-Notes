package main 

import "fmt"

func main(){
	// for loop 
	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n",i)
	}

	// while type for loop 
	var a , b = 1,5
	for a <= b {
		fmt.Println("a = ",a)
		a++ 
	}
}