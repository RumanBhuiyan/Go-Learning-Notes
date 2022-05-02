package main

import (
	"fmt"
	"strconv"
)

func main() {

	a := 5
	
	// Itoa => Integer to ASCII
	b := strconv.Itoa(a)
	fmt.Printf("value : %v Type : %[1]T \n",b)
	
	b = fmt.Sprintf("%v",a)
	fmt.Printf("value : %v Type : %[1]T \n",b)
	
	//  Atoi => ASCII to integer
	c,_ := strconv.Atoi(b)
	fmt.Printf("value : %v Type : %[1]T \n",c)

	// converting string to boolean
	str := "yes"
	var value bool
	
	switch str {
	     case "yes","true","1" :
				value = true
	     case "no","false","0" :
				value = false
	     default : 
			fmt.Println("string undefined")
	     
	}
	fmt.Println(value)	
}