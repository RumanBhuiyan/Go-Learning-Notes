package main

import (
	"fmt"
)

func main(){
	number := 2

	// if else ladder
	if number % 2 == 0{
		fmt.Println(number," is even")
	}else {
		fmt.Println(number," is odd")
	}

	// if - else if - else ladder, N.B => you can use else if without else when you need
	if number < 0 {
		fmt.Println("Negative")
	}else if number > 0 {
		fmt.Println("Positive")
	}else {
		fmt.Println("Zero")
	}

	// switch statement 
	switch number {
		case 0 : fmt.Println("Zero")
		case 1 : fmt.Println("One")
		case 2 : fmt.Println("Two")
		case 3 : fmt.Println("Three")
		case 4 : fmt.Println("Four")
		default : fmt.Println(number)
	}

	// fallthrough executes just next body statement of next case
	// in traditional c, next all cases body statements are executed if break isn't found or default case found
	switch number {
		case 0 : fmt.Println("Zero")
		case 1 : fmt.Println("One")
		case 2 : fmt.Println("Two")
					fallthrough
		case 3 : fmt.Println("Three")
					fallthrough
		case 4 : fmt.Println("Four")
					fallthrough
		default : fmt.Println(number)
	}

	// using variable short declaration in if-else,swith like in for loop
	if temp := 5; temp %2 == 0{
		fmt.Println("temp is even")
	}else {
		fmt.Println("temp is odd")
	}
	// temp is no longer availabe here it was alive just inside  above block
}