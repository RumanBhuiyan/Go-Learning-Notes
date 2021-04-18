package main

import "fmt"

func main(){
   
	var number int
	fmt.Printf("Enter the number: ")
	fmt.Scan(&number)
	fmt.Println("factorial of the number is ",getFactorial(number))
}

func getFactorial(number int) int {
   if (number == 1)  {
	   return 1 
	}
   return number * getFactorial(number -1 )
}