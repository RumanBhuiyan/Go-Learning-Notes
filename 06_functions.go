package main

import "fmt"

func main(){

  var number1, number2 float64
  fmt.Printf("Enter the first number: ")
  fmt.Scan(&number1)
  fmt.Printf("Enter the second number: ")
  fmt.Scan(&number2)

  fmt.Println("Summation: ",sum(int(number1),int(number2)))

  var product , quotient float64 = divisionResult(number1,number2)
  
  fmt.Println("Product :",product)
  fmt.Println("Quotient : ",quotient)
}

// function of returning one statement
func sum(a,b int ) int {
   return a+b 
}
// function of returning two statement
func divisionResult(a,b float64) (float64, float64) {
	return a*b , a / b ;
}