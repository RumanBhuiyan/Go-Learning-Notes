package main 

import "fmt"

func main(){
	// if - else if -else ladder 
	var number int 
	fmt.Printf("Enter the number: ")
	fmt.Scan(&number)

	if (number == 0){
		fmt.Println("Number is zero")
	}else if (number > 0){
		fmt.Println("Number is Positive")
	}else {
		fmt.Println("Number is Negative")
	}

	// switch statement 
	var another_number int = 2
	switch(another_number){
	case 0 :
		fmt.Println("zero")
		break 
	case 1 :
		fmt.Println("One")
		break 
	default :
	    fmt.Println("Undefined")
	}
}