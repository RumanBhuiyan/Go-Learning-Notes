package main

import "fmt"

func main(){
	var pointer *int 

	var keep int = 5;
	pointer = &keep  

	fmt.Println("Address: ",pointer,"value : ",*pointer)
}