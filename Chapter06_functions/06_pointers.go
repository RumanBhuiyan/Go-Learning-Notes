package main

import (
	"fmt"
)

func main(){
	/* &a -> address of a, & is called referencing operator
	 *&a -> value of the address of &a where * is called de-referencing operator  */

	a , b := 4,5
	fmt.Println("Before swaping  a : ",*&a," b : ",*&b)

	swap(&a, &b)
	fmt.Println("After swaping 	 a : ",a," b : ",b)

	// using new , new(data_type) returns a pointer of that data type 
	keep := new(int32)
	*keep = 5
	fmt.Println("Address of keep : ",&keep," value : ",*keep)
}

func swap(x , y *int) {
	temp := x
	*x = *y
	*y = *temp
}