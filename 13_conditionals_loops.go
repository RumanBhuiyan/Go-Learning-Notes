package main

import "fmt"

func main(){
	
	number := 23

	if number %2==0 {
		fmt.Println("Even number")
	}else {
		fmt.Println("Odd number")
	}

	// traditional for loop
	for i:=0 ; i<= 10 ; i++ {
		fmt.Println(i)
	}

	numbers := []int {1,3,5,}

	// as we aren't using index inside loop body then ignore it 
	// in place of index replace _
	for _,value := range numbers{
		fmt.Println(value)
	}

	// while style for loop
	i := 1
	for i<5{
		fmt.Println("Ruman")
		i++
	}

}

// parenthesis starting from newline is n't allowed in Go because it replace  ; 
//  after everyline automatically if there is no parenthesis. So below lines will
// be treated like func main() ; so correspoding lines are n't understandable by Go
// func main()
// {
// 	fmt.Println("Hello world")
// }