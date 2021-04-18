package main

import "fmt"

func main(){
	// process: 01 reading int number 
	var first_number int 
	fmt.Printf("Enter the first number: ")
	fmt.Scanf("%d",&first_number)

	// process: 02 reading int number 
	var second_number int 
	fmt.Printf("Enter Second number: ")
	fmt.Scan(&second_number)

	fmt.Println("summation : ",first_number+second_number)

	// reading string
	var country_name string 
	fmt.Printf("Enter country name : ")
	fmt.Scan(&country_name)
	fmt.Println("Your country is ",country_name)
}