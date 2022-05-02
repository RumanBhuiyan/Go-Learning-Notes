package main

import (
	"fmt"
)

func main() {
	
	a := 5
	
	fmt.Printf("type of a : %T value : %[1]v \n",a)
	
	/* %<width> => if width is less than required width then this width is neglected,
	 		if width is greater, then extra width is used for left padding,
			if width is negative ,then extra width is used for right padding
	
	 %<width>.(dot)<precision> => precision means how many digits will come after . and
				   width = <digits_before_dot>+1(for dot)+<digits_after_dot> */
	
	price := 1.0/3
	
	fmt.Println(price)
	fmt.Printf("price :%.4v\n",price)
	fmt.Printf("price :%9.4v\n",price)
}
