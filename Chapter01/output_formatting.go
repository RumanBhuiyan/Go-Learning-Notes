package main

import "fmt"

func main(){
	a := 5

	fmt.Print("value of a : " + fmt.Sprint(a))
	fmt.Print("\nvalue of a : " ,a)
	
	fmt.Println("\nvalue of a : " + fmt.Sprint(a))
	fmt.Println("value of a : ", a)
	
	fmt.Printf("\nvalue of a : " + fmt.Sprint(a))
	fmt.Printf("\nvalue of a : %v", a)

	// Printf() has an extra feature to offer where you can control width of outputs also
	// %+positive_width is used for left padding
	// %-negative width is used for right padding
	fmt.Printf("\nvalue of a : \t%15v", a)
	fmt.Printf("\nvalue of a : \t%-15v", a)
}